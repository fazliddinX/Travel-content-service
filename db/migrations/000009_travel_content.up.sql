CREATE TABLE messages (
                          id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                          sender_id UUID NOT NULL ,
                          recipient_id UUID NOT NULL,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

