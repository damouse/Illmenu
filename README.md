export GOOGLE_APPLICATION_CREDENTIALS="$(pwd)/token.json"

## Processing

Greyscale and resize images to reduce size. 

## OCR

Try out ABBYY. If a lot faster than tesseract it makes the image-upload time worth it. 

Name searching is only affected if it can detect bolded terms to google on the fly. May want to fall back to all-search. 

## All-Search

Start with text from menu
Search text for matches, return matched dishes
    - Fuzzy matching better
    - How does the text search work? Hard.

## Searching

Each dish has a *lastSearched* date.
If lastSearched == nil or lastSearched + 30 days < today, google image search

