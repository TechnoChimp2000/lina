package lina

import "errors"
// linear algebra package for matrix multiplication, vectors etc ...

// TODO: Define Matrix as an interface, and then implement its methods (
// TODO: Parallelize operations where possible - this part has to be done smartly. For small matrices, parallelization possibly leads to a 'slowdown', since the time spent on the overhead is longer than the benefit you gain from paralellization


// Vector is defined as a row vector
type Vector []float64

type Matrix []Vector

// 1.METHODS ON TYPE MATRIX

// 1.1 Basic methods
func ( m *Matrix ) GetRows() (rows int) {
	rows = len(*m)
	return
}

func ( m *Matrix ) GetColumns() (columns int) {
	columns = len((*m)[0])
	return
}

func ( m *Matrix ) GetSize() (elements int) {
	elements = m.GetRows() * m.GetColumns()
	return
}

func ( m *Matrix ) SetElement( row, column int, value float64) {
	(*m)[row][column] = value
}

// 1.2 Advanced Methods
func ( m *Matrix ) Transpose() {

	// create the new matrix and populate it with the new vectors
	r := make(Matrix, m.GetColumns())
	for i, _ := range r {
		r[i] = make(Vector, m.GetRows())
	}

	// populate the new matrix with the results
	for indexRow, row := range r {
		for indexColumn, _ := range row {
			r[indexRow][indexColumn] = (*m)[indexColumn][indexRow]
		}
	}

	// put the new Matrix under the memory address of the old Matrix -- ie we are modifying it
	*m = r
}

// 2. STANDALONE FUNCTIONS
func MatrixMultiplication( m1, m2 Matrix ) (r Matrix, err error) {

	// validate the correct matrix size!
	if m1.GetColumns() != m2.GetRows() {
		str := "error: number of rows in the first matrix does not equal with the number of columns in the second matrix"
		return nil, errors.New(str)
	}

	// create the result matrix (currently empty)
	r = make(Matrix, m1.GetRows())
	for i, _ := range r {
		r[i] = make(Vector, m2.GetColumns())
	}

	// populate the result matrix
	for indexRow, row := range m1 {
		for i:=0; i < m2.GetColumns(); i++ {
			for indexElement, element := range row {
				r[indexRow][i] += element*m2[indexElement][i]
			}
		}
	}
	return r, err
}
// matrix of equal sizes are multiplied element by element
func DotMultiply(m1, m2 Matrix) (r Matrix, err error) {

	// validate both matrices to confirm they are of equal sizes
	if (m1.GetRows() != m2.GetRows()) {
		str := "error: two matrices have to be of the same size. number of rows mismatch"
		return nil, errors.New(str)
	}

	if (m1.GetColumns() != m2.GetColumns() ) {
		str := "error: two matrices have to be of the same size. number of columns mismatch"
		return nil, errors.New(str)
	}



	// create result matrix
	r = make(Matrix, m1.GetRows())

	for i, _ := range r {
		r[i] = make(Vector, m1.GetColumns())
	}

	// multiply stuff
	for indexRow, row := range m1 {
		for indexColumn, _ := range row {
			r[indexRow][indexColumn] = m1[indexRow][indexColumn] * m2[indexRow][indexColumn]
		}
	}

	return r, err
}

// TODO: complete these methods here
func MatrixAdd(m1,m2 Matrix) (r Matrix, err error) {
	return
}

func MatrixSubstract(m1,m2 Matrix) (r Matrix, err error) {
	return
}


// 3. ERROR HANDLING
type errorString struct {
	s string
}

func (e *errorString) Error() string  {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}
