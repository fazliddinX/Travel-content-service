CREATE TABLE itineraries (
                             id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                             title VARCHAR(200) NOT NULL,
                             description TEXT,
                             start_date DATE NOT NULL,
                             end_date DATE NOT NULL,
                             author_id UUID NOT NULL,
                             likes_count INTEGER DEFAULT 0,
                             comments_count INTEGER DEFAULT 0,
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
