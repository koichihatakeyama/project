INSERT INTO users (
    name,
    email,
    created_at,
    updated_at
) VALUES (
    /* user.Name */'test',
    /* user.Email */'test@example.com',
    /* user.CreatedAt */CURRENT_TIMESTAMP,
    /* user.UpdatedAt */CURRENT_TIMESTAMP
)
RETURNING id
