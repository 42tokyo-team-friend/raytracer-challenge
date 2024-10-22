package goraytracer

import "math"

type Mat4x4 struct {
	data [4][4]float64
}

func NewIdentity() *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func (left *Mat4x4) Equals(right *Mat4x4) bool {
	flag := true

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if math.Abs(left.data[i][j] - right.data[i][j]) > Epsilon {
				flag = false
			}
		}
	}

	return flag
}

func (left *Mat4x4) Mul(right *Mat4x4) *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{
				left.data[0][0] * right.data[0][0] + left.data[0][1] * right.data[1][0] + left.data[0][2] * right.data[2][0] + left.data[0][3] * right.data[3][0],
				left.data[0][0] * right.data[0][1] + left.data[0][1] * right.data[1][1] + left.data[0][2] * right.data[2][1] + left.data[0][3] * right.data[3][1],
				left.data[0][0] * right.data[0][2] + left.data[0][1] * right.data[1][2] + left.data[0][2] * right.data[2][2] + left.data[0][3] * right.data[3][2],
				left.data[0][0] * right.data[0][3] + left.data[0][1] * right.data[1][3] + left.data[0][2] * right.data[2][3] + left.data[0][3] * right.data[3][3],
			},
			{
				left.data[1][0] * right.data[0][0] + left.data[1][1] * right.data[1][0] + left.data[1][2] * right.data[2][0] + left.data[1][3] * right.data[3][0],
				left.data[1][0] * right.data[0][1] + left.data[1][1] * right.data[1][1] + left.data[1][2] * right.data[2][1] + left.data[1][3] * right.data[3][1],
				left.data[1][0] * right.data[0][2] + left.data[1][1] * right.data[1][2] + left.data[1][2] * right.data[2][2] + left.data[1][3] * right.data[3][2],
				left.data[1][0] * right.data[0][3] + left.data[1][1] * right.data[1][3] + left.data[1][2] * right.data[2][3] + left.data[1][3] * right.data[3][3],
			},
			{
				left.data[2][0] * right.data[0][0] + left.data[2][1] * right.data[1][0] + left.data[2][2] * right.data[2][0] + left.data[2][3] * right.data[3][0],
				left.data[2][0] * right.data[0][1] + left.data[2][1] * right.data[1][1] + left.data[2][2] * right.data[2][1] + left.data[2][3] * right.data[3][1],
				left.data[2][0] * right.data[0][2] + left.data[2][1] * right.data[1][2] + left.data[2][2] * right.data[2][2] + left.data[2][3] * right.data[3][2],
				left.data[2][0] * right.data[0][3] + left.data[2][1] * right.data[1][3] + left.data[2][2] * right.data[2][3] + left.data[2][3] * right.data[3][3],
			},
			{
				left.data[3][0] * right.data[0][0] + left.data[3][1] * right.data[1][0] + left.data[3][2] * right.data[2][0] + left.data[3][3] * right.data[3][0],
				left.data[3][0] * right.data[0][1] + left.data[3][1] * right.data[1][1] + left.data[3][2] * right.data[2][1] + left.data[3][3] * right.data[3][1],
				left.data[3][0] * right.data[0][2] + left.data[3][1] * right.data[1][2] + left.data[3][2] * right.data[2][2] + left.data[3][3] * right.data[3][2],
				left.data[3][0] * right.data[0][3] + left.data[3][1] * right.data[1][3] + left.data[3][2] * right.data[2][3] + left.data[3][3] * right.data[3][3],
			},
		},
	}
}

func (m *Mat4x4) T() *Mat4x4 {
	return &Mat4x4{
		[4][4]float64{
			{m.data[0][0], m.data[1][0], m.data[2][0], m.data[3][0]},
			{m.data[0][1], m.data[1][1], m.data[2][1], m.data[3][1]},
			{m.data[0][2], m.data[1][2], m.data[2][2], m.data[3][2]},
			{m.data[0][3], m.data[1][3], m.data[2][3], m.data[3][3]},
		},
	}
}

func (m *Mat4x4) Inv() *Mat4x4 {
	dst := [4][4]float64{}
    /* Compute adjoint: */

    dst[0][0] = m.data[1][1]*m.data[2][2]*m.data[3][3] - m.data[1][1]*m.data[2][3]*m.data[3][2] -
		m.data[2][1]*m.data[1][2]*m.data[3][3] + m.data[2][1]*m.data[1][3]*m.data[3][2] +
		m.data[3][1]*m.data[1][2]*m.data[2][3] - m.data[3][1]*m.data[1][3]*m.data[2][2]

	dst[0][1] = -m.data[0][1]*m.data[2][2]*m.data[3][3] + m.data[0][1]*m.data[2][3]*m.data[3][2] +
		m.data[2][1]*m.data[0][2]*m.data[3][3] - m.data[2][1]*m.data[0][3]*m.data[3][2] -
		m.data[3][1]*m.data[0][2]*m.data[2][3] + m.data[3][1]*m.data[0][3]*m.data[2][2]

	dst[0][2] = m.data[0][1]*m.data[1][2]*m.data[3][3] - m.data[0][1]*m.data[1][3]*m.data[3][2] -
		m.data[1][1]*m.data[0][2]*m.data[3][3] + m.data[1][1]*m.data[0][3]*m.data[3][2] +
		m.data[3][1]*m.data[0][2]*m.data[1][3] - m.data[3][1]*m.data[0][3]*m.data[1][2]

	dst[0][3] = -m.data[0][1]*m.data[1][2]*m.data[2][3] + m.data[0][1]*m.data[1][3]*m.data[2][2] +
		m.data[1][1]*m.data[0][2]*m.data[2][3] - m.data[1][1]*m.data[0][3]*m.data[2][2] -
		m.data[2][1]*m.data[0][2]*m.data[1][3] + m.data[2][1]*m.data[0][3]*m.data[1][2]

	dst[1][0] = -m.data[1][0]*m.data[2][2]*m.data[3][3] + m.data[1][0]*m.data[2][3]*m.data[3][2] +
		m.data[2][0]*m.data[1][2]*m.data[3][3] - m.data[2][0]*m.data[1][3]*m.data[3][2] -
		m.data[3][0]*m.data[1][2]*m.data[2][3] + m.data[3][0]*m.data[1][3]*m.data[2][2]

	dst[1][1] = m.data[0][0]*m.data[2][2]*m.data[3][3] - m.data[0][0]*m.data[2][3]*m.data[3][2] -
		m.data[2][0]*m.data[0][2]*m.data[3][3] + m.data[2][0]*m.data[0][3]*m.data[3][2] +
		m.data[3][0]*m.data[0][2]*m.data[2][3] - m.data[3][0]*m.data[0][3]*m.data[2][2]

	dst[1][2] = -m.data[0][0]*m.data[1][2]*m.data[3][3] + m.data[0][0]*m.data[1][3]*m.data[3][2] +
		m.data[1][0]*m.data[0][2]*m.data[3][3] - m.data[1][0]*m.data[0][3]*m.data[3][2] -
		m.data[3][0]*m.data[0][2]*m.data[1][3] + m.data[3][0]*m.data[0][3]*m.data[1][2]

	dst[1][3] = m.data[0][0]*m.data[1][2]*m.data[2][3] - m.data[0][0]*m.data[1][3]*m.data[2][2] -
		m.data[1][0]*m.data[0][2]*m.data[2][3] + m.data[1][0]*m.data[0][3]*m.data[2][2] +
		m.data[2][0]*m.data[0][2]*m.data[1][3] - m.data[2][0]*m.data[0][3]*m.data[1][2]

	dst[2][0] = m.data[1][0]*m.data[2][1]*m.data[3][3] - m.data[1][0]*m.data[2][3]*m.data[3][1] -
		m.data[2][0]*m.data[1][1]*m.data[3][3] + m.data[2][0]*m.data[1][3]*m.data[3][1] +
		m.data[3][0]*m.data[1][1]*m.data[2][3] - m.data[3][0]*m.data[1][3]*m.data[2][1]

	dst[2][1] = -m.data[0][0]*m.data[2][1]*m.data[3][3] + m.data[0][0]*m.data[2][3]*m.data[3][1] +
		m.data[2][0]*m.data[0][1]*m.data[3][3] - m.data[2][0]*m.data[0][3]*m.data[3][1] -
		m.data[3][0]*m.data[0][1]*m.data[2][3] + m.data[3][0]*m.data[0][3]*m.data[2][1]

	dst[2][2] = m.data[0][0]*m.data[1][1]*m.data[3][3] - m.data[0][0]*m.data[1][3]*m.data[3][1] -
		m.data[1][0]*m.data[0][1]*m.data[3][3] + m.data[1][0]*m.data[0][3]*m.data[3][1] +
		m.data[3][0]*m.data[0][1]*m.data[1][3] - m.data[3][0]*m.data[0][3]*m.data[1][1]

	dst[2][3] = -m.data[0][0]*m.data[1][1]*m.data[2][3] + m.data[0][0]*m.data[1][3]*m.data[2][1] +
		m.data[1][0]*m.data[0][1]*m.data[2][3] - m.data[1][0]*m.data[0][3]*m.data[2][1] -
		m.data[2][0]*m.data[0][1]*m.data[1][3] + m.data[2][0]*m.data[0][3]*m.data[1][1]

	dst[3][0] = -m.data[1][0]*m.data[2][1]*m.data[3][2] + m.data[1][0]*m.data[2][2]*m.data[3][1] +
		m.data[2][0]*m.data[1][1]*m.data[3][2] - m.data[2][0]*m.data[1][2]*m.data[3][1] -
		m.data[3][0]*m.data[1][1]*m.data[2][2] + m.data[3][0]*m.data[1][2]*m.data[2][1]

	dst[3][1] = m.data[0][0]*m.data[2][1]*m.data[3][2] - m.data[0][0]*m.data[2][2]*m.data[3][1] -
		m.data[2][0]*m.data[0][1]*m.data[3][2] + m.data[2][0]*m.data[0][2]*m.data[3][1] +
		m.data[3][0]*m.data[0][1]*m.data[2][2] - m.data[3][0]*m.data[0][2]*m.data[2][1]

	dst[3][2] = -m.data[0][0]*m.data[1][1]*m.data[3][2] + m.data[0][0]*m.data[1][2]*m.data[3][1] +
		m.data[1][0]*m.data[0][1]*m.data[3][2] - m.data[1][0]*m.data[0][2]*m.data[3][1] -
		m.data[3][0]*m.data[0][1]*m.data[1][2] + m.data[3][0]*m.data[0][2]*m.data[1][1]

	dst[3][3] = m.data[0][0]*m.data[1][1]*m.data[2][2] - m.data[0][0]*m.data[1][2]*m.data[2][1] -
		m.data[1][0]*m.data[0][1]*m.data[2][2] + m.data[1][0]*m.data[0][2]*m.data[2][1] +
		m.data[2][0]*m.data[0][1]*m.data[1][2] - m.data[2][0]*m.data[0][2]*m.data[1][1]

	// Compute determinant
	det := m.data[0][0]*dst[0][0] + m.data[0][1]*dst[1][0] + m.data[0][2]*dst[2][0] + m.data[0][3]*dst[3][0]


    /* Multiply adjoint with reciprocal of determinant: */

    det = 1.0 / det;

    dst[0][0] *= det;
    dst[0][1] *= det;
    dst[0][2] *= det;
    dst[0][3] *= det;
    dst[1][0] *= det;
    dst[1][1] *= det;
    dst[1][2] *= det;
    dst[1][3] *= det;
    dst[2][0] *= det;
    dst[2][1] *= det;
    dst[2][2] *= det;
    dst[2][3] *= det;
    dst[3][0] *= det;
    dst[3][1] *= det;
    dst[3][2] *= det;
    dst[3][3] *= det;

    return &Mat4x4{
        dst,
    }
}

func (m *Mat4x4) MulTup(t Tuple) Tuple {
    return Tuple{
        m.data[0][0] * t.x + m.data[0][1] * t.y + m.data[0][2] * t.z + m.data[0][3] * t.w,
        m.data[1][0] * t.x + m.data[1][1] * t.y + m.data[1][2] * t.z + m.data[1][3] * t.w,
        m.data[2][0] * t.x + m.data[2][1] * t.y + m.data[2][2] * t.z + m.data[2][3] * t.w,
        m.data[3][0] * t.x + m.data[3][1] * t.y + m.data[3][2] * t.z + m.data[3][3] * t.w,
    }
}