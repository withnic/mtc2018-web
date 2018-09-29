// Package domains contains the types.
package domains

// Code generated by yo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// LikeSummaryServer represents a row from 'LikeSummaryServers'.
type LikeSummaryServer struct {
	Second    int64     `spanner:"Second" json:"Second"`       // Second
	SessionID int64     `spanner:"SessionId" json:"SessionId"` // SessionId
	ServerID  string    `spanner:"ServerId" json:"ServerId"`   // ServerId
	Likes     int64     `spanner:"Likes" json:"Likes"`         // Likes
	CreatedAt time.Time `spanner:"CreatedAt" json:"CreatedAt"` // CreatedAt
}

func LikeSummaryServerPrimaryKeys() []string {
	return []string{
		"Second",
		"SessionId",
		"ServerId",
	}
}

func LikeSummaryServerColumns() []string {
	return []string{
		"Second",
		"SessionId",
		"ServerId",
		"Likes",
		"CreatedAt",
	}
}

func (lss *LikeSummaryServer) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "Second":
			ret = append(ret, &lss.Second)
		case "SessionId":
			ret = append(ret, &lss.SessionID)
		case "ServerId":
			ret = append(ret, &lss.ServerID)
		case "Likes":
			ret = append(ret, &lss.Likes)
		case "CreatedAt":
			ret = append(ret, &lss.CreatedAt)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (lss *LikeSummaryServer) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "Second":
			ret = append(ret, lss.Second)
		case "SessionId":
			ret = append(ret, lss.SessionID)
		case "ServerId":
			ret = append(ret, lss.ServerID)
		case "Likes":
			ret = append(ret, lss.Likes)
		case "CreatedAt":
			ret = append(ret, lss.CreatedAt)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newLikeSummaryServer_Decoder returns a decoder which reads a row from *spanner.Row
// into LikeSummaryServer. The decoder is not goroutine-safe. Don't use it concurrently.
func newLikeSummaryServer_Decoder(cols []string) func(*spanner.Row) (*LikeSummaryServer, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*LikeSummaryServer, error) {
		var lss LikeSummaryServer
		ptrs, err := lss.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &lss, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (lss *LikeSummaryServer) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("LikeSummaryServers", LikeSummaryServerColumns(), []interface{}{
		lss.Second, lss.SessionID, lss.ServerID, lss.Likes, lss.CreatedAt,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (lss *LikeSummaryServer) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("LikeSummaryServers", LikeSummaryServerColumns(), []interface{}{
		lss.Second, lss.SessionID, lss.ServerID, lss.Likes, lss.CreatedAt,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (lss *LikeSummaryServer) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("LikeSummaryServers", LikeSummaryServerColumns(), []interface{}{
		lss.Second, lss.SessionID, lss.ServerID, lss.Likes, lss.CreatedAt,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (lss *LikeSummaryServer) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, LikeSummaryServerPrimaryKeys()...)

	values, err := lss.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "LikeSummaryServer.UpdateColumns", "LikeSummaryServers", err)
	}

	return spanner.Update("LikeSummaryServers", colsWithPKeys, values), nil
}

// FindLikeSummaryServer gets a LikeSummaryServer by primary key
func FindLikeSummaryServer(ctx context.Context, db YORODB, second int64, sessionID int64, serverID string) (*LikeSummaryServer, error) {
	key := spanner.Key{second, sessionID, serverID}
	row, err := db.ReadRow(ctx, "LikeSummaryServers", key, LikeSummaryServerColumns())
	if err != nil {
		return nil, newError("FindLikeSummaryServer", "LikeSummaryServers", err)
	}

	decoder := newLikeSummaryServer_Decoder(LikeSummaryServerColumns())
	lss, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindLikeSummaryServer", "LikeSummaryServers", err)
	}

	return lss, nil
}

// Delete deletes the LikeSummaryServer from the database.
func (lss *LikeSummaryServer) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := lss.columnsToValues(LikeSummaryServerPrimaryKeys())
	return spanner.Delete("LikeSummaryServers", spanner.Key(values))
}
