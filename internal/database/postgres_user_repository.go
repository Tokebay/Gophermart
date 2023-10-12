// package database

// import "github.com/Tokebay/yandex-diplom/internal/models"

// type PostgresUserRepository struct {
// 	db *sqlx.DB
// }

// func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
// 	return &PostgresUserRepository{db: db}
// }

// func (r *PostgresUserRepository) CreateUser(user *models.User) error {
// 	_, err := r.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, user.Password)
// 	return err
// }

// func (r *PostgresUserRepository) GetUserByLogin(login string) (*models.User, error) {
// 	user := &models.User{}
// 	err := r.db.Get(user, "SELECT * FROM users WHERE login = $1", login)
// 	return user, err
// }
