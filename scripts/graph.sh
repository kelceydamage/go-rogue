./main 2> graph.dot
dot -Tpng graph.dot -o graph.png
open graph.png