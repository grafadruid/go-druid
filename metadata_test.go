package druid

import (
	_ "embed"
	"errors"
	"strings"
	"time"

	"github.com/h2oai/go-druid/builder/query"
)

func (c *Client) metadata(options ...metadataOption) *metadataService {
	return newMetadataService(c, options...)
}

type count struct {
	Cnt int `json:"cnt"`
}

type metadataOptions struct {
	tickerDuration time.Duration
	awaitTimeout   time.Duration
}

type metadataOption func(*metadataOptions)

// metadataService is a service that runs druid metadata requests using druid SQL API.
// NOTE: for internal tests use only, not representing official druid API.
type metadataService struct {
	client         *Client
	tickerDuration time.Duration
	awaitTimeout   time.Duration
}

func newMetadataService(client *Client, options ...metadataOption) *metadataService {
	opts := &metadataOptions{
		tickerDuration: 500 * time.Millisecond,
		awaitTimeout:   180 * time.Second,
	}
	for _, opt := range options {
		opt(opts)
	}
	md := &metadataService{
		client:         client,
		tickerDuration: opts.tickerDuration,
		awaitTimeout:   opts.awaitTimeout,
	}
	return md
}

func WithMetadataQueryTicker(duration time.Duration) metadataOption {
	return func(opts *metadataOptions) {
		opts.tickerDuration = duration
	}
}

func WithMetadataQueryTimeout(timeout time.Duration) metadataOption {
	return func(opts *metadataOptions) {
		opts.awaitTimeout = timeout
	}
}

//go:embed sql/datasource_available.sql
var datasourceAvailableQuery string

func fillDataSourceName(in string, ds string) string {
	return strings.Replace(in, "${{ datasource }}", ds, 1)
}

// awaitDataSourceAvailable awaits for a datasource to be visible in druid table listing.
// NOTE: for internal tests use only, not representing official druid API.
func (md *metadataService) awaitDataSourceAvailable(dataSourceName string) error {
	ticker := time.NewTicker(md.tickerDuration)
	defer ticker.Stop()
	afterTimeout := time.After(md.awaitTimeout)
	q := query.
		NewSQL().
		SetQuery(datasourceAvailableQuery).
		SetParameters([]query.SQLParameter{query.NewSQLParameter("VARCHAR", dataSourceName)})
	for {
		select {
		case <-ticker.C:
			var res []count
			_, err := md.client.Query().Execute(q, &res)
			if err != nil {
				return err
			}
			if len(res) >= 1 && res[0].Cnt == 1 {
				return nil
			}
		case <-afterTimeout:
			return errors.New("awaitDataSourceAvailable timeout")
		}
	}
}

//go:embed sql/datasource_records.sql
var datasourceRecordsQuery string

// awaitRecordsCount awaits for specific recordsCount in a given datasource.
// NOTE: not safe and intended for internal tests use only. Not representing official druid API.
func (md *metadataService) awaitRecordsCount(dataSourceName string, recordsCount int) error {
	ticker := time.NewTicker(md.tickerDuration)
	defer ticker.Stop()
	q := query.NewSQL()
	q.SetQuery(fillDataSourceName(datasourceRecordsQuery, dataSourceName))
	afterTimeout := time.After(md.awaitTimeout)
	for {
		select {
		case <-ticker.C:
			var res []count
			_, err := md.client.Query().Execute(q, &res)
			if err != nil {
				return err
			}

			if len(res) >= 1 && res[0].Cnt == recordsCount {
				return nil
			}
		case <-afterTimeout:
			return errors.New("awaitRecordsCount timeout")
		}
	}
}
