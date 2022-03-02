package pkg

func (m *Morakab) CreateBook(title, author string) error {
	if _, err := m.DB.Exec(`INSERT INTO books (title, writer) VALUES ($1, $2)`, title, author); err != nil {
		return err
	}
	return nil
}
