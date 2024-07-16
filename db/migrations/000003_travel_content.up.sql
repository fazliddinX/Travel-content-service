CREATE TABLE comments (
                          id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                          content TEXT NOT NULL,
                          author_id UUID NOT NULL ,
                          story_id UUID REFERENCES stories(id),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
