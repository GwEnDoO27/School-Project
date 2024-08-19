objects = src/main.o src/system.o src/auth.o 
LIBS =  -lsqlite3 

atm: $(objects)
		clang -o atm $(objects) $(LIBS)

clean:
		rm -f $(objects) atm