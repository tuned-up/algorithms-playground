package main

import (
	"fmt"
)

func main()  {
	awg := arrayWithGaps{}
	awg.Set(0,2, 0)
	awg.Set(2,4, 6)
	awg.Set(2,5, 7)
	awg.Set(3,1, 1)
	awg.Print()
}

type arrayWithGaps struct {
	width, height int
	start *arrayWithGapsRow
}

type arrayWithGapsRow struct {
	rowNumber int
	nextRow   *arrayWithGapsRow
	start     *arrayWithGapsEntry
}

type arrayWithGapsEntry struct {
	entryNumber int
	value int
	nextEntry *arrayWithGapsEntry
}

func (a *arrayWithGaps) Print() {
	var targetEntry *arrayWithGapsEntry
	targetRow := a.start
	for i := 0; i < a.height + 1; i++ {
		if targetRow != nil && targetRow.rowNumber == i {
			targetEntry = targetRow.start
			targetRow = targetRow.nextRow
		}
		for j := 0; j < a.width + 1; j++ {
			if targetEntry != nil && targetEntry.entryNumber == j {
				fmt.Printf(" %d ", targetEntry.value)
				targetEntry = targetEntry.nextEntry
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
}

func (a *arrayWithGaps) Find(row, column int) (int, bool) {
	currentRow := a.start
	for currentRow != nil && currentRow.rowNumber <= row {
		if currentRow.rowNumber == row {
			currentEntry := currentRow.start
			for currentEntry != nil && currentEntry.entryNumber <= column {
				if currentEntry.entryNumber == column {
					return currentEntry.value, true
				}

				currentEntry = currentEntry.nextEntry
			}

			return 0, false
		}

		currentRow = currentRow.nextRow
	}

	return 0, false
}

func (a *arrayWithGaps) Set(row, column, value int) (*arrayWithGapsRow, *arrayWithGapsEntry) {
	if a.height < row {
		a.height = row
	}
	if a.width < column {
		a.width = column
	}
	if a.start == nil {
		targetRow := &arrayWithGapsRow{rowNumber: row}
		a.start = targetRow
	}
	currentRow := a.start

	for currentRow.nextRow != nil && currentRow.rowNumber < row {
		currentRow = currentRow.nextRow
	}

	var targetRow *arrayWithGapsRow
	if currentRow.rowNumber == row {
		targetRow = currentRow
	} else if currentRow.nextRow == nil {
		targetRow = &arrayWithGapsRow{rowNumber: row}
		currentRow.nextRow = targetRow
	} else if currentRow.nextRow.rowNumber > row {
		targetRow = &arrayWithGapsRow{rowNumber: row, nextRow: currentRow.nextRow}
		currentRow.nextRow = targetRow
	}

	var targetEntry *arrayWithGapsEntry
	if targetRow.start == nil {
		targetEntry = &arrayWithGapsEntry{entryNumber: column, value: value}
		targetRow.start = targetEntry
	}

	currentEntry := targetRow.start
	for currentEntry.nextEntry != nil && currentEntry.entryNumber < column {
		currentEntry = currentEntry.nextEntry
	}
	if currentEntry.entryNumber == column {
		currentEntry.value = value
	} else if currentEntry.nextEntry == nil {
		targetEntry = &arrayWithGapsEntry{entryNumber: column, value: value}
		currentEntry.nextEntry = targetEntry
	} else if currentEntry.nextEntry.entryNumber == column {
		currentEntry.nextEntry.value = value
	} else if currentEntry.nextEntry.entryNumber > column {
		targetEntry = &arrayWithGapsEntry{entryNumber: column, nextEntry: currentEntry.nextEntry}
		currentEntry.nextEntry = targetEntry
	}
	
	return targetRow, targetEntry
}

func (a *arrayWithGaps) Del(row, column int) (int, bool) {
	currentRow := a.start

	for currentRow.nextRow != nil && currentRow.nextRow.rowNumber < row {
		currentRow = currentRow.nextRow
	}

	if currentRow.nextRow == nil || currentRow.rowNumber > row {
		return 0, false
	}

	currentEntry := currentRow.start
	if currentEntry == nil {
		return 0, false
	}

	for currentEntry.nextEntry != nil && currentEntry.nextEntry.entryNumber < column {
		currentEntry = currentEntry.nextEntry
	}
	
	if currentEntry.nextEntry == nil || currentEntry.nextEntry.entryNumber > column {
		return 0, false
	}

	value := currentEntry.nextEntry.value
	afterNext := currentEntry.nextEntry.nextEntry
	currentEntry.nextEntry = afterNext

	return value, true
}

