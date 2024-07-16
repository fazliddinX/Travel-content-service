CREATE TABLE stories (
                         id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                         title VARCHAR(200) NOT NULL,
                         content TEXT NOT NULL,
                         location VARCHAR(100),
                         author_id UUID NOT NULL ,
                         likes_count INTEGER DEFAULT 0,
                         comments_count INTEGER DEFAULT 0,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
