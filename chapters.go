package bookstack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Chapter struct {
	ID          int       `json:"id,omitempty"`
	BookID      int       `json:"book_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Slug        string    `json:"slug,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	CreatedAt   string    `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedBy   int       `json:"created_by,omitempty"`
	UpdatedBy   int       `json:"updated_by,omitempty"`
	OwnedBy     int       `json:"owned_by,omitempty"`
}

type ChapterDetailed struct {
	ID          int           `json:"id,omitempty"`
	BookID      int           `json:"book_id,omitempty"`
	Slug        string        `json:"slug,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Priority    int           `json:"priority,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	CreatedBy   CreatedBy     `json:"created_by,omitempty"`
	UpdatedBy   UpdatedBy     `json:"updated_by,omitempty"`
	OwnedBy     OwnedBy       `json:"owned_by,omitempty"`
	Tags        []Tag         `json:"tags,omitempty"`
	Pages       []ChapterPage `json:"pages,omitempty"`
}

type ChapterPage struct {
	ID            int       `json:"id,omitempty"`
	BookID        int       `json:"book_id,omitempty"`
	ChapterID     int       `json:"chapter_id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Slug          string    `json:"slug,omitempty"`
	Priority      int       `json:"priority,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	CreatedBy     int       `json:"created_by,omitempty"`
	UpdatedBy     int       `json:"updated_by,omitempty"`
	Draft         bool      `json:"draft,omitempty"`
	RevisionCount int       `json:"revision_count,omitempty"`
	Template      bool      `json:"template,omitempty"`
}

type ChapterParams struct {
	BookID          int         `json:"book_id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Description     string      `json:"description,omitempty"`
	DescriptionHtml string      `json:"description_html,omitempty"`
	Tags            []TagParams `json:"tags,omitempty"`
}

func (bp ChapterParams) Form() (string, io.Reader, error) {

	r, err := json.Marshal(bp)
	if err != nil {
		return "", nil, err
	}

	return appJSON, bytes.NewReader(r), nil
}

// ListChapters will return the chapters that match the given params.
func (b *Bookstack) ListChapters(ctx context.Context, params *QueryParams) ([]Chapter, error) {

	resp, err := b.request(ctx, http.MethodGet, params.String("/chapters"), blank{})
	if err != nil {
		return nil, err
	}

	return ParseMultiple[[]Chapter](resp)
}

// GetChapter will return a single chapter that matches id.
func (b *Bookstack) GetChapter(ctx context.Context, id int) (ChapterDetailed, error) {

	resp, err := b.request(ctx, http.MethodGet, fmt.Sprintf("/chapters/%d", id), blank{})
	if err != nil {
		return ChapterDetailed{}, err
	}

	return ParseSingle[ChapterDetailed](resp)

}

// CreateChapter will create a chapter according to the given params.
func (b *Bookstack) CreateChapter(ctx context.Context, params ChapterParams) (Chapter, error) {

	resp, err := b.request(ctx, http.MethodPost, "/chapters", params)
	if err != nil {
		return Chapter{}, err
	}

	return ParseSingle[Chapter](resp)
}

// UpdateChapter will update a chapter with the given params.
func (b *Bookstack) UpdateChapter(ctx context.Context, id int, params ChapterParams) (Chapter, error) {

	resp, err := b.request(ctx, http.MethodPut, fmt.Sprintf("/chapters/%d", id), params)
	if err != nil {
		return Chapter{}, err
	}

	return ParseSingle[Chapter](resp)
}

// DeleteChapter will delete a chapter with the given id.
func (b *Bookstack) DeleteChapter(ctx context.Context, id int) (bool, error) {

	if _, err := b.request(ctx, http.MethodDelete, fmt.Sprintf("/chapters/%d", id), blank{}); err != nil {
		return false, err
	}

	return true, nil
}

// ExportChapterHTML will return a chapter in HTML format.
func (b *Bookstack) ExportChapterHTML(ctx context.Context, id int) (io.Reader, error) {

	data, err := b.request(ctx, http.MethodGet, fmt.Sprintf("/chapters/%d/export/html", id), blank{})
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil

}

// ExportChapterPDF will return a chapter in PDF format.
func (b *Bookstack) ExportChapterPDF(ctx context.Context, id int) (io.Reader, error) {

	data, err := b.request(ctx, http.MethodGet, fmt.Sprintf("/chapters/%d/export/pdf", id), blank{})
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil

}

// ExportChapterMarkdown will return a chapter in Markdown format.
func (b *Bookstack) ExportChapterMarkdown(ctx context.Context, id int) (io.Reader, error) {

	data, err := b.request(ctx, http.MethodGet, fmt.Sprintf("/chapters/%d/export/markdown", id), blank{})
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil

}

// ExportChapterPlaintext will return a chapter in Plaintext format.
func (b *Bookstack) ExportChapterPlaintext(ctx context.Context, id int) (io.Reader, error) {

	data, err := b.request(ctx, http.MethodGet, fmt.Sprintf("/chapters/%d/export/plaintext", id), blank{})
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil

}
