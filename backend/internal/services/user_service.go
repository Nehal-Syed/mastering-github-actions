package services

import (
    "database/sql"
    "mastering-actions/internal/models"
)

type UserService struct {
    db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{db: db}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
    result, err := s.db.Exec(
        "INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
        req.Name, req.Email, req.Age,
    )
    if err != nil {
        return nil, err
    }

    id, _ := result.LastInsertId()
    return s.GetUserByID(int(id))
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
    var user models.User
    err := s.db.QueryRow(
        "SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
    
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    rows, err := s.db.Query("SELECT id, name, email, age, created_at, updated_at FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (s *UserService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
    query := "UPDATE users SET "
    args := []interface{}{}
    
    if req.Name != "" {
        query += "name = ?, "
        args = append(args, req.Name)
    }
    if req.Email != "" {
        query += "email = ?, "
        args = append(args, req.Email)
    }
    if req.Age > 0 {
        query += "age = ?, "
        args = append(args, req.Age)
    }
    
    query = query[:len(query)-2] + " WHERE id = ?"
    args = append(args, id)
    
    _, err := s.db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    
    return s.GetUserByID(id)
}

func (s *UserService) DeleteUser(id int) error {
    _, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}