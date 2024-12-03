package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Section struct {
	From, To int
}

type RangeDelta struct {
	From, To int
	Delta    int
}

type (
	Block             []*RangeDelta
	InstructionBlocks [7]*Block
)

// genInstructionsBlock converts a list of instructions from the input
// into a sorted list of ranges with the delta for this range
func genInstructionBlock(instructionLines []string) *Block {
	var block Block

	// Convert Input to Ranges
	for _, instructionLine := range instructionLines {
		instructionData := ExtractNumbers(instructionLine)

		block = append(block, &RangeDelta{
			From:  instructionData[1],
			To:    instructionData[1] + instructionData[2] - 1,
			Delta: instructionData[0] - instructionData[1],
		})
	}

	// Sort Ranges
	slices.SortFunc(block, func(i, j *RangeDelta) int {
		if i.From < j.From {
			return -1
		}
		return 1
	})

	return &block
}

// transformMappingTable applies one instruction block to a mapping table to
// get a new mapping table (i.e. a table that maps input ranges to delta values
// which incorporate the instructions from this block
func transformMappingTable(mappingTable *Block, instructionBlock *Block) *Block {
	newTable := Block{}

	for _, mappingEntry := range *mappingTable {

		// remainder saves how much of the current mappingEntry we have not yet processed
		// we add the current delta sum so that the range is correct for the input to
		// this stage. It will be subtracted later again.
		remainder := &Section{
			From: mappingEntry.From + mappingEntry.Delta,
			To:   mappingEntry.To + mappingEntry.Delta,
		}

		// Now apply instructions until we have processed the whole mappingEntry
		for _, instruction := range *instructionBlock {
			if instruction.To < remainder.From {
				// instruction range ends before this mapping entry starts, so skip
				continue
			}

			if instruction.From > remainder.To {
				// instruction range starts after this mapping entry ends, so we are done
				break
			}

			if instruction.From > remainder.From {
				// The beginning of this mapping entry stays at the same delta
				newTable = append(newTable, &RangeDelta{
					From:  remainder.From - mappingEntry.Delta,
					To:    instruction.From - mappingEntry.Delta - 1,
					Delta: mappingEntry.Delta,
				})
				remainder.From = instruction.From
			}

			// now apply the additional delta from the instruction
			newTable = append(newTable, &RangeDelta{
				From:  remainder.From - mappingEntry.Delta,
				To:    min(instruction.To, remainder.To) - mappingEntry.Delta,
				Delta: mappingEntry.Delta + instruction.Delta,
			})

			// if we end after the end of the instruction range, we leave the rest for the
			// next instruction
			if instruction.To < remainder.To {
				remainder.From = instruction.To + 1
			} else {
				// otherwise we have no remainder
				remainder = nil
				break
			}
		}

		// If we still have a remainder, append it with unmodified delta
		if remainder != nil {
			newTable = append(newTable, &RangeDelta{
				From:  remainder.From - mappingEntry.Delta,
				To:    remainder.To - mappingEntry.Delta,
				Delta: mappingEntry.Delta,
			})
		}

	}

	return &newTable
}

func calc(input *Input) (int, int) {
	resultPart1 := math.MaxInt
	resultPart2 := math.MaxInt

	// Convert Instructions
	textBlocks := input.TextBlocks()
	instructionBlocks := InstructionBlocks{}
	for block := 1; block < 8; block++ {
		instructionBlocks[block-1] = genInstructionBlock(textBlocks[block][1:])
	}

	// The Mapping Table maps input ranges to delta values
	// (which are to be added in this range) to get from input to output.
	// We start with a single entry convering the full integer range,
	// which will be split up by applying the instruction blocks.
	mappingTable := &Block{
		&RangeDelta{From: 0, To: math.MaxInt, Delta: 0},
	}

	// Now apply instruction blocks to get final input to output mapping
	for _, instructionBlock := range instructionBlocks {
		mappingTable = transformMappingTable(mappingTable, instructionBlock)
	}

	// part 1
	// this could also be solved by the optimized method of part 2, but after all
	// the hours of writing the optimized code I want to keep the brute force solution
	// at least for this :) This is still somewhat optimized because we use the
	// computed mapping table.
	seeds := ExtractNumbers(textBlocks[0][0])
	for _, seed := range seeds {
		for _, mappingEntry := range *mappingTable {
			if seed >= mappingEntry.From && seed <= mappingEntry.To {
				seed += mappingEntry.Delta
				break
			}
		}

		if seed < resultPart1 {
			resultPart1 = seed
		}
	}

	// part2

	// sort the mapping table so that the mappings that yield the lowest results are first
	slices.SortFunc(*mappingTable, func(i, j *RangeDelta) int {
		if i.From+i.Delta < j.From+j.Delta {
			return -1
		}
		return 1
	})

	// Go through the mapping table until we find one that applies to one of our inputs,
	// then we have the result
search:
	for _, mappingEntry := range *mappingTable {
		for num := 0; num < len(seeds); num += 2 {
			firstSeed := seeds[num]
			numSeeds := seeds[num+1]

			// this mapping entry does not apply to this seed range
			if firstSeed > mappingEntry.To || firstSeed+numSeeds < mappingEntry.From {
				continue
			}

			// if our seed range starts before the start of the mapping entry
			// we use the start of the mapping entry to get the location
			// otherwise we get the location from the first seed of the range
			resultPart2 = max(firstSeed, mappingEntry.From) + mappingEntry.Delta
			break search
		}
	}

	return resultPart1, resultPart2
}

func main() {
	//Run("Sample", "sample1.txt", calc)
	Run("Main", "./input.txt", calc)
}

//-----------------------------------------------------------------------------
// run.go

// Run runs the given calcFunction on the given input file
func Run(text string, fileName string, calcFunction InputCalcFunction) {
	//_ = pp.Print // just to keep this module in the project

	InputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Errorf("%s: %w", text, err))
		return
	}

	var lines []string

	scanner := bufio.NewScanner(InputFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	res1, res2 := calcFunction(NewInput(lines))
	fmt.Printf("%s: %d, %d\n", text, res1, res2)
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// tools.go

// SliceMemberOrEmptyString returns the member of a slice at the given index,
// or an empty string if the index is out of bounds
func SliceMemberOrEmptyString(slice []string, index int) string {
	if index < len(slice) {
		return slice[index]
	}
	return ""
}

// Atoi converts a string to an int, ignoring errors (return zero instead)
func Atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func ExtractRegexps(s, expr string) []string {
	re := regexp.MustCompile(expr)
	return re.FindAllString(s, -1)
}

// ExtractNumbers extracts all numbers from a string
func ExtractNumbers(s string) []int {
	var res []int

	for _, match := range ExtractRegexps(s, `\d+`) {
		res = append(res, Atoi(match))
	}
	return res
}

func ExtractDigits(s string) []int {
	var res []int

	for _, match := range ExtractRegexps(s, `\d`) {
		res = append(res, Atoi(match))
	}
	return res
}

func RegexpSubmatchAsInt(s, expr string) int {
	re := regexp.MustCompile(expr)
	match := re.FindStringSubmatch(s)
	return Atoi(SliceMemberOrEmptyString(match, 1))
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// line.go

// Line represents a line of input
type Line struct {
	LineNo int
	Data   string
}

// FindObjects returns a list of objects that match the given regular expression on this Line
func (l *Line) FindObjects(re string) []*Object {
	var objects []*Object
	matcher := regexp.MustCompile(re)

	matches := matcher.FindAllStringIndex(l.Data, -1)
	for _, match := range matches {
		objects = append(objects, &Object{
			Line:  l,
			left:  match[0],
			right: match[1] - 1,
		})
	}

	return objects
}

// ReplaceText replaces all occurrences of the given string with the given replacement
// on this Line. Note: length of find and replace must be the same
func (l *Line) ReplaceText(find, replace string) {
	if len(find) != len(replace) {
		return
	}
	l.Data = strings.ReplaceAll(l.Data, find, replace)
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// object.go

// Object represents an object found on a line
type Object struct {
	Line        *Line
	left, right int
}

// String returns the object as a string
func (o *Object) String() string {
	return o.Line.Data[o.left : o.right+1]
}

// Int returns the object as an int
func (o *Object) Int() int {
	return Atoi(o.String())
}

// Adjacent returns true if the given object is adjacent to this object
func (o *Object) Adjacent(other *Object) bool {
	if other.right < o.left-1 || other.left > o.right+1 {
		return false
	}

	if other.Line.LineNo < o.Line.LineNo-1 || other.Line.LineNo > o.Line.LineNo+1 {
		return false
	}

	return true
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// input.go

// Input represents the input data
type Input struct {
	Lines []*Line
}

// InputCalcFunction is the function signature for the calculation function
type InputCalcFunction func(i *Input) (int, int)

// NewInput creates a new Input object from the given input lines
func NewInput(inputLines []string) *Input {
	i := &Input{}

	for lineNo, line := range inputLines {
		i.Lines = append(i.Lines, &Line{
			LineNo: lineNo,
			Data:   line,
		})
	}

	return i
}

// FindObjects returns a list of objects that match the given regular expression
func (i *Input) FindObjects(re string) []*Object {
	var objects []*Object

	for _, line := range i.Lines {
		lineObjects := line.FindObjects(re)
		objects = append(objects, lineObjects...)
	}

	return objects
}

// TextBlocks returns the input data as a list of blocks (separated by empty lines in input)
func (i *Input) TextBlocks() [][]string {
	var blocks [][]string
	var block []string

	for _, line := range i.Lines {
		if line.Data == "" {
			blocks = append(blocks, block)
			block = []string{}
			continue
		}
		block = append(block, line.Data)
	}
	blocks = append(blocks, block)

	return blocks
}

//-----------------------------------------------------------------------------
