Question:
  Can we create a model that describes how 1-dimensional cellular automata calculate?

Details:
  I know cellular automata can calculate the solution to certain problems. These calculations can be described using a particle-model. But can we use these particles to program a cellular atomata by hand?

Sub-questions:
  - Can we deduce the rules that will produce particles given the local neighborhood?
  - How can we know how particles will interact? I.e. can we construct the particle catalog?
  - How can we connect what we know from the answer to the two questions above to construct a high-fitness solution?

Strategy:
  - Try replicating the density classification task, but using adaptive neighborhood search instead of a genetic algorithm.

Cheatsheet for rules:

111 110 101 100 011 010 001 000   // Radius = 1

110100011


11110 11101 11010 10100 01000 10001 00011 00111 01111  
  1     0     0     1     0     0     1     1     1
  30    29    26    20    8     17    3     7     15
