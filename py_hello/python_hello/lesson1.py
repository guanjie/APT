"""Python ppt tutorial"""

class myfunctions:
    def exp():
        return 0

class atom:
    def __init__(self, atno, x, y, z):
        self.atno = atno
        self.__position = (x, y, z)
    def __repr__(self):
        return '%d %10.4f %10.4f %10.4f' % (self.atno, self.__position[0],
                self.__position[1], self.__position[2])
    def symbol(self):
        return self.atno

    def getposition(self):
        return self.__position
    def setposition(self, x, y, z):
        self.__position = (x, y, z) #typecheck first!
    def tranlate(self, x, y, z):
        x0, y0, z0 = self.__position
        self.__position = (x0+x, y0+y, z0+z)

class molecule:
    def __init__(self, name='Generic'):
        self.name = name
        self.atomlist = []
    def addatom(self, atom):
        self.atomlist.append(atom)
    def __getitem__(self, index):
        return self.atomlist[index]
    def __repr__(self):
        str = 'This is a molecule named %s\n' % self.name
        str = str + 'It has %d atoms\n' % len(self.atomlist)
        for atom in self.atomlist:
            str = str + `atom` + '\n'
        return str

class qm_molecule(molecule):
    def addbasis(self):
        self.basis = []
        for atom in self.atomlist:
            self.basis = add_bf(atom, self.basis)
    def __repr__(self):
        str = 'QM Rules! \n'
        for atom in self.atomlist:
            str = str + `atom` + '\n'
        return str


