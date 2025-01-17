package main


func encrypt (text, key string) {
  key = removeSpaces(toUpperCase(key))
  numColumns := len(key)
  numRows := (len(text) + numColumns - 1 + numColumns - 1) / numColumns

  table := make([][]rune, numRows)
  for i := range table {
    table[i] = make([]rune, numColumns)
  }

  // filling the transpositioj table …
  index := 0
  for i := 0; i < numRows; i++ {
    for j := 0; j < numColumns; j++ {
      if index < len(text) {
        table[i][j] = rune(text[index])
        index++
      } else {
        table[i][j] = 'X'
      }
    }
  }

  perm := sortedKeyIndices(key)
	fmt.Println(perm)

	var encryptedText string
	for _, colIndex := range perm {
		for _, row := range table {
			if colIndex < len(row) {
				encryptedText += string(row[colIndex])
			}
		}
	}
	return encryptedText	
}
