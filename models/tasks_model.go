package models

import (
	"database/sql"
	"log"
	"task-management-api/config"
	"task-management-api/entities"
)

func GetTasks() []entities.Tasks {
	rows, err := config.DB.Query(`SELECT * FROM tasks`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var tasks []entities.Tasks

	for rows.Next() {
		var task entities.Tasks
		err := rows.Scan(
			&task.ID,
			&task.TaskName,
			&task.TaskDesc,
			&task.CreatedBy,
			&task.UpdatedBy,
		)

		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func GetTaskByID(id int) (entities.Tasks, error) {
	var task entities.Tasks

	err := config.DB.QueryRow(`SELECT * FROM tasks WHERE ID = ?`, id).Scan(
		&task.ID,
		&task.TaskName,
		&task.TaskDesc,
		&task.CreatedBy,
		&task.UpdatedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No task found with ID %d", id)
			return task, err
		}
		log.Printf("Error fetching task with ID %d: %v", id, err)
		return task, err
	}

	log.Printf("Task found: %+v", task) // Logging task yang ditemukan
	return task, nil
}

func CreateTask(task entities.Tasks) error {
	// Query untuk insert data
	query := `INSERT INTO tasks (taskName, taskDesc, createdBy, updatedBy) VALUES (?, ?, ?, ?)`

	// Eksekusi query
	_, err := config.DB.Exec(query, task.TaskName, task.TaskDesc, task.CreatedBy, task.UpdatedBy)

	if err != nil {
		log.Printf("Error creating task: %v", err)
		return err
	}

	log.Println("Task created successfully..")
	return nil
}

func UpdateTask(id int, task entities.Tasks) error {
	// Query untuk update data
	query := `UPDATE tasks SET taskName = ?, taskDesc = ?, updatedBy = ? WHERE ID = ?`

	// Eksekusi query
	_, err := config.DB.Exec(query, task.TaskName, task.TaskDesc, task.UpdatedBy, id)

	if err != nil {
		log.Printf("Error updating task with ID %d: %v", id, err)
		return err
	}

	log.Printf("Task with ID %d updated successfully..", id)
	return nil
}

func DeleteTask(id int) error {
	// Query untuk delete data
	query := `DELETE FROM tasks WHERE ID = ?`

	// Eksekusi query
	_, err := config.DB.Exec(query, id)

	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", id, err)
		return err
	}

	log.Printf("Task with ID %d deleted successfully..", id)
	return nil
}
