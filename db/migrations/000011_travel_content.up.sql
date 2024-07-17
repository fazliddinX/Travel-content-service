CREATE TABLE itineraries_likes (
                             user_id UUID NOT NULL ,
                             itineraries_id UUID REFERENCES itineraries(id),
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);