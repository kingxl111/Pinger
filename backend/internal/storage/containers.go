package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kingxl111/Pinger/backend/internal/models"
)

var _ ContainerManager = (*ContainerManagerPG)(nil)

type ContainerManagerPG struct {
	db *DB
}

func NewContainersPG(db *DB) *ContainerManagerPG {
	return &ContainerManagerPG{db: db}
}

func (c *ContainerManagerPG) NewContainer(ctx context.Context, container models.ContainerPing) error {
	const op = "ContainerManager.NewContainer"
	tx, err := c.db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", err, op)
	}

	builder := sq.Insert(containerTable).
		PlaceholderFormat(sq.Dollar).
		Columns(containerTableIPColumn, containerTablePingTime, containerTableLastSuccessPing).
		Values(container.IP, container.PingTime, container.LastSuccessPingTime).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %s", err, op)
	}
	if err := tx.QueryRow(ctx, query, args...).Scan(&container.ID); err != nil {
		return fmt.Errorf("%w: %s", err, op)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", err, op)
	}
	return nil
}

func (c *ContainerManagerPG) GetContainers(ctx context.Context) ([]models.ContainerPing, error) {
	return nil, nil
}
func (c *ContainerManagerPG) GetContainer(ctx context.Context, containerID int) (models.ContainerPing, error) {
	return models.ContainerPing{}, nil
}
