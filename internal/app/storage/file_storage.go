package storage

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
)

type FileStorage struct {
	filePath string
}

func newFileStorage(storagePath string) (*FileStorage, error) {
	return &FileStorage{filePath: storagePath}, nil
}

func (stge *FileStorage) SaveLinkToStge(hash string, original []byte) error {
	id := uuid.New()

	record := URLRecord{
		ID:          id.String(),
		ShortURL:    hash,
		OriginalURL: string(original),
	}

	writer, err := newWriter(stge.filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer writer.close()

	if err := writer.writeRecord(&record); err != nil {
		log.Fatal(err)
	}

	return err
}

func (stge *FileStorage) GetLinkFromStge(hash []byte) ([]byte, error) {
	reader, err := newReader(stge.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.close()

	readEvent, err := reader.readRecord(string(hash))
	if err != nil {
		log.Fatal(err)
	}

	return []byte(readEvent.OriginalURL), nil
}

type URLRecord struct {
	ID          string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type storageWriter struct {
	file    *os.File
	encoder *json.Encoder
}

func newWriter(filename string) (*storageWriter, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &storageWriter{
		file,
		json.NewEncoder(file),
	}, nil
}

func (p *storageWriter) close() error {
	return p.file.Close()
}

func (p *storageWriter) writeRecord(rec *URLRecord) error {
	data, err := json.Marshal(&rec)
	if err != nil {
		return err
	}
	// добавляем перенос строки
	data = append(data, '\n')

	_, err = p.file.Write(data)
	return err
}

type storageReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func newReader(filename string) (*storageReader, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &storageReader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, err
}

func (c *storageReader) readRecord(hash string) (*URLRecord, error) {
	for c.scanner.Scan() {
		data := c.scanner.Bytes()
		record := &URLRecord{}
		err := json.Unmarshal(data, &record)
		if err != nil {
			return nil, err
		}

		if record.ShortURL == hash {
			return record, nil
		}
	}

	return nil, c.scanner.Err()
}

func (c *storageReader) close() error {
	return c.file.Close()
}
