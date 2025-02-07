## Memory Is Managed by the Runtime
- Memory is allocated and deallocated automatically
- The Go runtime is statically linked
- Memory is managed in the runtime in dedicated threads

## Memory Allocation
### Use make() or new() to initialize complex objects
- maps
- slices
- channels

### The new() function
- Allocates, but does not initialize memory
- Results is zeroed storage but returns a memory

### The make() function
- Allocates and initilizes memory
- Allocates non-zeroed storage and returns a memory address

### Creating nil Objects
- You must initialize complex objects before adding values
- Declaration without make() can cause a panic


