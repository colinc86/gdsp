# gDSP
[![GoDoc](https://godoc.org/github.com/colinc86/gdsp?status.svg)](https://godoc.org/github.com/colinc86/gdsp) [![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fcolinc86%2Fgdsp%2Fbadge&style=flat)](https://actions-badge.atrox.dev/colinc86/gdsp/goto)

A digital signal processing library for Go.

## Features
gDSP is a light-weight vector library with a focus on digital signal processing. It is not meant to replace other more robust linear algebra packages.

### Functions
- [x] Autoregressive model parameters using Burg's method
- [x] Autocorrelation
- [x] Convolution
- [x] Cross-correlation
- [x] Discrete Fourier transform
- [x] Fast Fourier transform
- [x] Extrapolation
- [x] 1-dimensional digital filter
- [x] Filter initialization function
- [x] IIR filter
- [ ] FIR filter
- [x] Interpolation
- [x] Gaussian lowpass filter
- [x] Normalization

### Windowing
- [x] Hann
- [x] Hamming
- [x] Nuttal

### Vectors
- [x] Real and complex support
- [x] Vector arithmetic
- [x] Padding functions

### Matrices
- [ ] Conjugate
- [ ] Determinant
- [ ] Transpose
