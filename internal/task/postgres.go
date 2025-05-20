package task

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (r *PostgresRepository) Create(task *Task) error {
	_, err := r.DB.ExecContext(context.Background(),
		`INSERT INTO tasks (pr_id, release_id, type, content) VALUES ($1, $2, $3, $4)`,
		task.PrId, task.ReleaseId, task.Type, task.Content)

	return err
}

func (r *PostgresRepository) ListByReleaseId(releaseId string) ([]Task, error) {
	rows, err := r.DB.QueryContext(context.Background(),
		`SELECT pr_id, release_id, type, content FROM tasks WHERE release_id = $1`,
		releaseId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		err := rows.Scan(&t.PrId, &t.ReleaseId, &t.Type, &t.Content)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)

	}

	return tasks, err
}

func (r *PostgresRepository) DeleteById(id int) error {
	_, err := r.DB.ExecContext(context.Background(),
		`DELETE FROM tasks WHERE id = $1`,
		id)

	return err
}
