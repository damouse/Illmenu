export GOOGLE_APPLICATION_CREDENTIALS="$(pwd)/token.json"

## Query

Client must send OCR text body
Search text for matches, return matched dishes
    - Fuzzy matching better


## Searching

Each dish has a *lastSearched* date.
If lastSearched == nil or lastSearched + 30 days < today, google image search
