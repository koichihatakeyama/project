package entity

import "time"

type User struct {
    ID        int64     `db:"id"`
    Name      string    `db:"name"`
    CreatedAt time.Time `db:"created_at"`
}
