CREATE TABLE travel_tips (
                             id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                             title VARCHAR(200) NOT NULL,
                             content TEXT NOT NULL,
                             category VARCHAR(50),
                             author_id UUID NOT NULL ,
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

