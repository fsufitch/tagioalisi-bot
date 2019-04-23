
def chunk_lines(lines: [str], max_chars=1900) -> [[str]]:
    chunks = []
    current_chunk = []
    current_chunk_len = 0
    for line in lines:
        if current_chunk_len + len(line) > max_chars:
            chunks.append(current_chunk)
            current_chunk = []
            current_chunk_len = 0
        current_chunk.append(line)
        current_chunk_len += len(line) + 1 # account for '\n'
    chunks.append(current_chunk)
    return chunks
