# Hyperpath routing in Golang

Just implementation of [Spiess, H. and Florian, M. (1989) "Optimal strategies: A new assignment model for transit networks"](https://doi.org/10.1016/0191-2615(89)90034-9) in Golang

## Algorithm

Here is copy of algorithm in MathJax (for the LaTeX see [spiess_floarian.tex](./spiess_floarian.tex)):

### Part 1: Find optimal strategy

1. **Initialization**
   - Set $u_r = 0$ for destination node
   - Set $u_i = \infty$ for all other nodes
   - Set $f_i = 0$ for all nodes
   - Initialize empty attractive set $\overline{A}$

2. **Label Setting**
   - For each link $a = (i,j)$ with minimum $u_j + c_a$
   - If $u_i \geq u_j + c_a$:
     * Update node label: $$u_i = \frac{f_i \cdot u_i + f_a \cdot (u_j + c_a)}{f_i + f_a}$$
     * Update frequency: $$f_i = f_i + f_a$$
     * Add to attractive set: $$\overline{A} = \overline{A} \cup \{a\}$$
    
### Part 2: Assign demand according to optimal strategy

1. **Initialization**
   - Set $V_i = g_i$ for all nodes

2. **Loading**
   - Process links in decreasing order of $u_j + c_a$
   - For attractive links $a \in \overline{A}$:
     * Calculate volume: $$v_a = \frac{f_a}{f_i}V_i$$
     * Update node volume: $$V_j = V_j + v_a$$


## How to use

* Get the package:
   ```shell
   go get github.com/lddl/go-hyperpaths
   ```

* Code (you can find it in [examples/paper](./examples/paper))
   ```go
   // @todo
   ```

## References
Spiess, H. and Florian, M. (1989) "Optimal strategies: A new assignment model for transit networks". Transportation Research Part B: Methodological, 23(2), 83-102. Available in: https://doi.org/10.1016/0191-2615(89)90034-9