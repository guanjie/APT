# Triangle Project Code.

# Triangle analyzes the lengths of the sides of a triangle
# (represented by a, b and c) and returns the type of triangle.
#
# It returns:
#   :equilateral  if all sides are equal
#   :isosceles    if exactly 2 sides are equal
#   :scalene      if no sides are equal
#
# The tests for this method can be found in
#   about_triangle_project.rb
# and
#   about_triangle_project_2.rb
#
def triangle(a, b, c)
  # WRITE THIS CODE
  # Raise the errors first
    if a <= 0 || b <= 0 || c <= 0
        raise TriangleError, "No equal less than 0 Error"
    elsif a + b <=c || a + c <= b || b + c <= a
        raise TriangleError, "Two sides summation less than 3rd Error"

        # Here comes the normal conditions
    elsif a == b  && a == c
        :equilateral
    elsif a == b || a == c || b == c
       :isosceles
    else
        :scalene
    end
end

# Error class used in part 2.  No need to change this code.
class TriangleError < StandardError
end
