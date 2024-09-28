package domains

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-shiori/shiori/internal/core"
	"github.com/go-shiori/shiori/internal/model"
	"github.com/sirupsen/logrus"
)

type ArchiverDomain struct {
	dataDir     string
	openAiKey   string
	openAiModel string
	logger      *logrus.Logger
}

func (d *ArchiverDomain) DownloadBookmarkArchive(book model.Bookmark) (*model.Bookmark, error) {
	content, contentType, err := core.DownloadBookmark(book.URL)
	if err != nil {
		return nil, fmt.Errorf("error downloading url: %s", err)
	}

	d.logger.Warnf("d.openAiKey %v", d.openAiKey)

	processRequest := core.ProcessRequest{
		DataDir:     d.dataDir,
		OpenAiKey:   d.openAiKey,
		OpenAiModel: d.openAiModel,
		Bookmark:    book,
		Content:     content,
		ContentType: contentType,
	}

	result, isFatalErr, err := core.ProcessBookmark(processRequest)
	content.Close()

	if err != nil && isFatalErr {
		return nil, fmt.Errorf("failed to process: %v", err)
	}

	return &result, nil
}

func (d *ArchiverDomain) GetBookmarkArchive(book model.Bookmark) error {
	archivePath := filepath.Join(d.dataDir, "archive", strconv.Itoa(book.ID))

	info, err := os.Stat(archivePath)
	if !os.IsNotExist(err) && !info.IsDir() {
		return fmt.Errorf("archive not found")
	}

	return nil
}

func NewArchiverDomain(logger *logrus.Logger, dataDir string, openAiKey string, openAiModel string) ArchiverDomain {
	return ArchiverDomain{
		dataDir:     dataDir,
		openAiKey:   openAiKey,
		openAiModel: openAiModel,
		logger:      logger,
	}
}
