package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/go-practice-01/internal/tasks"
)

func Load(path string) (tasks.State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return tasks.State{NextID: 1}, nil
		}
		return tasks.State{}, fmt.Errorf("чтение файла %s: %w", path, err)
	}
	var state tasks.State
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&state); err != nil {
		return tasks.State{}, fmt.Errorf("декодирование JSON в %s: %w", path, err)
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return tasks.State{}, fmt.Errorf("файл %s должен содержать один JSON объект", path)
	}

	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return tasks.State{}, fmt.Errorf("валидация состояния в %s: %w", path, err)
	}
	return state, nil
}

func Save(path string, state tasks.State) (err error) {
	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return fmt.Errorf("валидация состояния перед сохранением в %s: %w", path, err)
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("кодирование JSON ля %s: %w", path, err)
	}
	data = append(data, '\n')

	tmp, err := os.CreateTemp(filepath.Dir(path), ".tasks-*.tmp")
	if err != nil {
		return fmt.Errorf("создание временного файла %s: %w", path, err)
	}

	tmpName := tmp.Name()
	defer func() {
		if err != nil {
			_ = os.Remove(tmpName)
		}
	}()

	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("запись во временный файл для %s: %w", path, err)
	}

	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("синхронизация временного файла для %s: %w", path, err)
	}

	if err := tmp.Close(); err != nil {
		return fmt.Errorf("закрытие временного файла для %s: %w", path, err)
	}

	if err := os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("замена файла %s: %w", path, err)
	}

	return nil
}
