-- the original migration did not properly establish cascading constraints for the meme tables
-- this script does that

ALTER TABLE ONLY public.meme_names DROP CONSTRAINT IF EXISTS meme_names_meme_id_fkey CASCADE;
ALTER TABLE ONLY public.meme_names 
    ADD CONSTRAINT meme_names_meme_id_fkey FOREIGN KEY (meme_id) REFERENCES public.memes(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.meme_urls DROP CONSTRAINT IF EXISTS meme_urls_meme_id_fkey CASCADE;
ALTER TABLE ONLY public.meme_urls
    ADD CONSTRAINT meme_urls_meme_id_fkey FOREIGN KEY (meme_id) REFERENCES public.memes(id) ON DELETE CASCADE;