package python

//#include "Python.h"
//#include <stdlib.h>
//#include <string.h>
//int _gopy_PyObject_CheckBuffer(PyObject *obj) { return PyObject_CheckBuffer(obj); }
import "C"
//import "unsafe"
import "os"

// Py_buffer layer
type Py_buffer struct {
	ptr *C.Py_buffer
}

type PyBUF_Flag int

const (
	PyBUF_SIMPLE = PyBUF_Flag(C.PyBUF_SIMPLE)
	PyBUF_WRITABLE = PyBUF_Flag(C.PyBUF_WRITABLE)
	PyBUF_STRIDES = PyBUF_Flag(C.PyBUF_STRIDES)
	PyBUF_ND = PyBUF_Flag(C.PyBUF_ND)
	PyBUF_C_CONTIGUOUS = PyBUF_Flag(C.PyBUF_C_CONTIGUOUS)
	PyBUF_INDIRECT = PyBUF_Flag(C.PyBUF_INDIRECT)
	PyBUF_FORMAT = PyBUF_Flag(C.PyBUF_FORMAT)
	PyBUF_STRIDED = PyBUF_Flag(C.PyBUF_STRIDED)
	PyBUF_STRIDED_RO = PyBUF_Flag(C.PyBUF_STRIDED_RO)
	PyBUF_RECORDS = PyBUF_Flag(C.PyBUF_RECORDS)
	PyBUF_RECORDS_RO = PyBUF_Flag(C.PyBUF_RECORDS_RO)
	PyBUF_FULL = PyBUF_Flag(C.PyBUF_FULL)
	PyBUF_FULL_RO = PyBUF_Flag(C.PyBUF_FULL_RO)
	PyBUF_CONTIG = PyBUF_Flag(C.PyBUF_CONTIG)
	PyBUF_CONTIG_RO = PyBUF_Flag(C.PyBUF_CONTIG_RO)
)

/*
int PyObject_CheckBuffer(PyObject *obj)
Return 1 if obj supports the buffer interface otherwise 0.
*/
func PyObject_CheckBuffer(self *PyObject) bool {
	return int2bool(C._gopy_PyObject_CheckBuffer(topy(self)))
}

/*
int PyObject_GetBuffer(PyObject *obj, Py_buffer *view, int flags)
Export obj into a Py_buffer, view. These arguments must never be NULL. The flags argument is a bit field indicating what kind of buffer the caller is prepared to deal with and therefore what kind of buffer the exporter is allowed to return. The buffer interface allows for complicated memory sharing possibilities, but some caller may not be able to handle all the complexity but may want to see if the exporter will let them take a simpler view to its memory.

Some exporters may not be able to share memory in every possible way and may need to raise errors to signal to some consumers that something is just not possible. These errors should be a BufferError unless there is another error that is actually causing the problem. The exporter can use flags information to simplify how much of the Py_buffer structure is filled in with non-default values and/or raise an error if the object can’t support a simpler view of its memory.

0 is returned on success and -1 on error.

The following table gives possible values to the flags arguments.

Flag	Description
PyBUF_SIMPLE	This is the default flag state. The returned buffer may or may not have writable memory. The format of the data will be assumed to be unsigned bytes. This is a “stand-alone” flag constant. It never needs to be ‘|’d to the others. The exporter will raise an error if it cannot provide such a contiguous buffer of bytes.
PyBUF_WRITABLE	The returned buffer must be writable. If it is not writable, then raise an error.
PyBUF_STRIDES	This implies PyBUF_ND. The returned buffer must provide strides information (i.e. the strides cannot be NULL). This would be used when the consumer can handle strided, discontiguous arrays. Handling strides automatically assumes you can handle shape. The exporter can raise an error if a strided representation of the data is not possible (i.e. without the suboffsets).
PyBUF_ND	The returned buffer must provide shape information. The memory will be assumed C-style contiguous (last dimension varies the fastest). The exporter may raise an error if it cannot provide this kind of contiguous buffer. If this is not given then shape will be NULL.
PyBUF_C_CONTIGUOUS PyBUF_F_CONTIGUOUS PyBUF_ANY_CONTIGUOUS	These flags indicate that the contiguity returned buffer must be respectively, C-contiguous (last dimension varies the fastest), Fortran contiguous (first dimension varies the fastest) or either one. All of these flags imply PyBUF_STRIDES and guarantee that the strides buffer info structure will be filled in correctly.
PyBUF_INDIRECT	This flag indicates the returned buffer must have suboffsets information (which can be NULL if no suboffsets are needed). This can be used when the consumer can handle indirect array referencing implied by these suboffsets. This implies PyBUF_STRIDES.
PyBUF_FORMAT	The returned buffer must have true format information if this flag is provided. This would be used when the consumer is going to be checking for what ‘kind’ of data is actually stored. An exporter should always be able to provide this information if requested. If format is not explicitly requested then the format must be returned as NULL (which means 'B', or unsigned bytes)
PyBUF_STRIDED	This is equivalent to (PyBUF_STRIDES | PyBUF_WRITABLE).
PyBUF_STRIDED_RO	This is equivalent to (PyBUF_STRIDES).
PyBUF_RECORDS	This is equivalent to (PyBUF_STRIDES | PyBUF_FORMAT | PyBUF_WRITABLE).
PyBUF_RECORDS_RO	This is equivalent to (PyBUF_STRIDES | PyBUF_FORMAT).
PyBUF_FULL	This is equivalent to (PyBUF_INDIRECT | PyBUF_FORMAT | PyBUF_WRITABLE).
PyBUF_FULL_RO	This is equivalent to (PyBUF_INDIRECT | PyBUF_FORMAT).
PyBUF_CONTIG	This is equivalent to (PyBUF_ND | PyBUF_WRITABLE).
PyBUF_CONTIG_RO	This is equivalent to (PyBUF_ND).
*/
func PyObject_GetBuffer(self *PyObject, flags PyBUF_Flag) (buf *Py_buffer, err os.Error) {
	buf.ptr = &C.Py_buffer{}
	err = int2err(C.PyObject_GetBuffer(topy(self), buf.ptr, C.int(flags)))
	return
}
/*
void PyBuffer_Release(Py_buffer *view)
Release the buffer view. This should be called when the buffer is no longer being used as it may free memory from it.
*/
func PyBuffer_Release(self *Py_buffer) {
	C.PyBuffer_Release(self.ptr)
}

/*
Py_ssize_t PyBuffer_SizeFromFormat(const char *)
Return the implied ~Py_buffer.itemsize from the struct-stype ~Py_buffer.format.
*/
func PyBuffer_SizeFromFormat(self *Py_buffer) int {
	//FIXME
	panic("not implemented")
}

/*
int PyObject_CopyToObject(PyObject *obj, void *buf, Py_ssize_t len, char fortran)
Copy len bytes of data pointed to by the contiguous chunk of memory pointed to by buf into the buffer exported by obj. The buffer must of course be writable. Return 0 on success and return -1 and raise an error on failure. If the object does not have a writable buffer, then an error is raised. If fortran is 'F', then if the object is multi-dimensional, then the data will be copied into the array in Fortran-style (first dimension varies the fastest). If fortran is 'C', then the data will be copied into the array in C-style (last dimension varies the fastest). If fortran is 'A', then it does not matter and the copy will be made in whatever way is more efficient.
*/
func PyObject_CopyToObject(self *PyObject, buf []byte, fortran string) os.Error {
	/*
	c_buf := (*C.char)(unsafe.Pointer(&buf[0]))
	c_for := C.char(fortran[0])

	py_self := self.ptr
	//defer C.free(unsafe.Pointer(c_for))
	err := C.PyObject_CopyToObject(py_self, c_buf, C.Py_ssize_t(len(buf)), c_for)
	 */
	//FIXME
	panic("not implemented")
}

/*
int PyBuffer_IsContiguous(Py_buffer *view, char fortran)
Return 1 if the memory defined by the view is C-style (fortran is 'C') or Fortran-style (fortran is 'F') contiguous or either one (fortran is 'A'). Return 0 otherwise.
*/
func PyBuffer_IsContiguous(self *Py_buffer, fortran string) bool {
	c_fortran := C.char(fortran[0])
	return int2bool(C.PyBuffer_IsContiguous(self.ptr, c_fortran))
}

/*
void PyBuffer_FillContiguousStrides(int ndim, Py_ssize_t *shape, Py_ssize_t *strides, Py_ssize_t itemsize, char fortran)
Fill the strides array with byte-strides of a contiguous (C-style if fortran is 'C' or Fortran-style if fortran is 'F' array of the given shape with the given number of bytes per element.
*/
func PyBuffer_FillContiguousStrides(ndim int, shape, strides []int, itemsize int, fortran string) {
	//FIXME
	panic("not implemented")
}

/*
int PyBuffer_FillInfo(Py_buffer *view, PyObject *obj, void *buf, Py_ssize_t len, int readonly, int infoflags)
Fill in a buffer-info structure, view, correctly for an exporter that can only share a contiguous chunk of memory of “unsigned bytes” of the given length. Return 0 on success and -1 (with raising an error) on error.
*/
func PyBuffer_FillInfo(self *PyObject, buf []byte, readonly bool, infoflags int) (buffer *Py_buffer, err os.Error) {
	//FIXME
	panic("not implemented")
}

// EOF

