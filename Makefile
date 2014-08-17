.PHONY: clean

CXX_FLAGS= -std=c++11 -stdlib=libstdc++
LD_FLAGS = -L. -I . -I/usr/local/include -lvw -stdlib=libstdc++

TARGET=go-vw

#$(TARGET): libstubs.a
# go build .

$(TARGET): stubs.so
	go build -ldflags -L. .

stubs.so: stubs.cpp
	g++ $(LD_FLAGS) `go env GOGCCFLAGS` $(CXX_FLAGS) -shared -o stubs.so stubs.cpp

#libstubs.a: stubs.o
	#ar r $@ $^

#%.o: %.cpp
	#g++ $(CXX_FLAGS) -I/usr/local/include -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
