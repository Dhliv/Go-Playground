package CP

import "math"

const epsilon float64 = 0.0000000000000000000000001

type simplex struct {
	matrix          [][]float64
	posMaxImprove   int
	posMinAvailable int
	answer          []float64
	justInteger     bool
}

func (s *simplex) findMaxImprove() bool {
	var haveMaxPositive bool = false
	var maxi float64 = 0

	for i := 1; i < len(s.matrix[0]); i++ {
		if maxi < s.matrix[0][i] && math.Abs(maxi-s.matrix[0][i]) > epsilon {
			maxi = s.matrix[0][i]
			haveMaxPositive = true
			s.posMaxImprove = i
		}
	}

	return haveMaxPositive
}

func (s *simplex) findMinAvailable() bool {
	var mini, aux float64 = math.MaxFloat64, 0
	var haveMinAvailable bool = false

	for i := 1; i < len(s.matrix); i++ { // we don't explore the function to maximize.
		if s.matrix[i][s.posMaxImprove] == 0.0 { // We don't want any division by zero.
			continue
		}

		aux = s.matrix[i][0]/(-s.matrix[i][s.posMaxImprove]) + epsilon
		if math.Abs(aux-epsilon) < epsilon {
			aux = 0
		}
		if aux > epsilon && mini > aux && math.Abs(mini-aux) > epsilon {
			mini = aux
			haveMinAvailable = true
			s.posMinAvailable = i
		}
	}

	if mini == 0 {
		haveMinAvailable = false
		s.matrix[0][s.posMaxImprove] = 0
	}

	return haveMinAvailable
}

func (s *simplex) improveAnswer() {
	var div float64 = -s.matrix[s.posMinAvailable][s.posMaxImprove]
	for i := 0; i < len(s.matrix); i++ {
		s.matrix[s.posMaxImprove][i] = s.matrix[s.posMinAvailable][i]/div + epsilon

		if math.Abs(s.matrix[s.posMaxImprove][i]-epsilon) < epsilon {
			s.matrix[s.posMaxImprove][i] = 0
		}

		s.matrix[s.posMinAvailable][i] = 0
	}

	s.matrix[s.posMaxImprove][s.posMaxImprove] = 0
	// Until here, we isolated the var s.posMaxImprove from equation s.posMinAvailable.

	// Now we have to replace the var s.posMaxImprove in every single equation:

	for i := 0; i < len(s.matrix); i++ {
		// we don't check the same equation 'cause its value of itself it's zero, so it won't change.
		for j := 0; j < len(s.matrix); j++ {
			s.matrix[i][j] += s.matrix[s.posMaxImprove][j]*s.matrix[i][s.posMaxImprove] + epsilon

			if math.Abs(s.matrix[i][j]-epsilon) < epsilon {
				s.matrix[i][j] = 0
			}
		}
		s.matrix[i][s.posMaxImprove] = 0
	}
}

// TODO we need more theory
func (s *simplex) getAnswer() {
	s.answer = make([]float64, len(s.matrix))
	s.answer[0] = s.matrix[0][0]
}

func Simplex(a, b float64) {
	s := simplex{
		matrix: [][]float64{
			// all the variables in order, left to right are the same as up to down.
			// Column doesn't include the var for the function.
			// first column is the result of every equation that we have.
			{0, 0, 0, 1, 1, 1},
			{a, 0, 0, -3, -2, -1},
			{b, 0, 0, -1, -2, -3},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
	}

	for s.findMaxImprove() {
		if !s.findMinAvailable() {
			continue
		}

		s.improveAnswer()
		//Printf("matrix: %v\ns.posMaxImprove = %v and s.posMinavailable = %v\n", s.matrix, s.posMaxImprove, s.posMinAvailable)
	}

	s.getAnswer()
	Print(s.answer[0])
}
