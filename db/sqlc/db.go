// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createPostStmt, err = db.PrepareContext(ctx, createPost); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePost: %w", err)
	}
	if q.deletePostStmt, err = db.PrepareContext(ctx, deletePost); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePost: %w", err)
	}
	if q.getPostByIdStmt, err = db.PrepareContext(ctx, getPostById); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostById: %w", err)
	}
	if q.listPostsStmt, err = db.PrepareContext(ctx, listPosts); err != nil {
		return nil, fmt.Errorf("error preparing query ListPosts: %w", err)
	}
	if q.updatePostStmt, err = db.PrepareContext(ctx, updatePost); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePost: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createPostStmt != nil {
		if cerr := q.createPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPostStmt: %w", cerr)
		}
	}
	if q.deletePostStmt != nil {
		if cerr := q.deletePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePostStmt: %w", cerr)
		}
	}
	if q.getPostByIdStmt != nil {
		if cerr := q.getPostByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostByIdStmt: %w", cerr)
		}
	}
	if q.listPostsStmt != nil {
		if cerr := q.listPostsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPostsStmt: %w", cerr)
		}
	}
	if q.updatePostStmt != nil {
		if cerr := q.updatePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePostStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db              DBTX
	tx              *sql.Tx
	createPostStmt  *sql.Stmt
	deletePostStmt  *sql.Stmt
	getPostByIdStmt *sql.Stmt
	listPostsStmt   *sql.Stmt
	updatePostStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:              tx,
		tx:              tx,
		createPostStmt:  q.createPostStmt,
		deletePostStmt:  q.deletePostStmt,
		getPostByIdStmt: q.getPostByIdStmt,
		listPostsStmt:   q.listPostsStmt,
		updatePostStmt:  q.updatePostStmt,
	}
}
