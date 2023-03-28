CREATE TABLE public.books (
    id bigint NOT NULL,
    title character varying(100) NOT NULL,
    author character varying(100) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
