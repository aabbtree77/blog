# Tensor Algebra

[Donn G. Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) reveals how to parse mysterious phrases such as (i) "tensor fields contain irreducible representations of the Lorentz group", (ii) "antisymmetric two-tensor field can carry two particle triplets", and (iii) "a free-field equation is nothing but an invariant record of which components are superfluous". I will add here a few notes that might help one to decipher his paper.

An archetypical problem that ChatGPT cannot solve: Given the [four-vector](https://en.wikipedia.org/wiki/Four-vector) $k_{\mu}$ and [the metric tensor](https://en.wikipedia.org/wiki/Metric_tensor) $g_{\mu\nu}$, write down the most general dimensionless tensor ${T_{\mu\nu}}^{\rho \sigma}$ symmetric under the permutations of the covariant indices $(\mu, \nu)$, and also symmetric w.r.t. the permutations of contravariant indices $(\rho, \sigma)$. It should be a sum of linearly independent terms, each with a manifest symmetry, and at most fourth order in $k_{\mu}$.

[Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) jumps into the answer, which is a linear combination of

$$
\begin{align}
X_{1} & = \frac{1}{4} {g_{(\mu}}^{(\rho} {g_{\nu)}}^{\sigma)}\,,\\
X_{2} & = g_{\mu \nu} g^{\rho \sigma}\,, \\
X_{3} & = \frac{1}{k^2} g_{\mu \nu} k^{\rho} k^{\sigma}\,, \\
X_{4} & = \frac{1}{k^2} k_{\mu} k_{\nu} g^{\rho \sigma}\,, \\
X_{5} & = \frac{1}{k^4} k_{\mu} k_{\nu} k^{\rho} k^{\sigma}\,, \\
X_{6} & = \frac{1}{k^2} k_{(\mu} {g_{\nu)}}^{(\rho} k^{\sigma)}\,.
\end{align}
$$

Here the division by 4 of the first basis element is not important, but it would turn $X_{1}$ into an index symmetrization operator if one acted with it on any two-index tensor.

The logic is clear. One groups the combinations of primitives according to their distinct combinatorial types. Linear independence is verified by choosing a few random non-zero elements of a tensor.

One can view tensors as matrices by appending each element with an indicator basis:

$$
\begin{align}
e^{\mu}\otimes e^{\nu}\otimes e_{\rho}\otimes e_{\sigma}\,,
\end{align}
$$

with the implied summation over repeated indices. Thus, each $X_{i}$ can be seen as the sum of 256 matrices of size 16x16. Here one can assume that the upper/lower index corresponds to the row/column vector, resp.

[The Kronecker product](https://en.wikipedia.org/wiki/Kronecker_product) has an interesting property:

$$
\begin{align}
a^{i}\otimes b^{j} &\neq b^{j}\otimes a^{i}\,,\\
a^{i}\otimes b_{j} &= b_{j}\otimes a^{i}\,.
\end{align}
$$

Along with the associativity, this reveals the tensor index positioning ambiguity:

$$
X{^i}{^j}{_k} = X{^i}{_k}{^j} = X{_k}{^i}{^j}.
$$

It is tempting to write $X^{ij}_{k}$, but then one cannot always raise or lower an index uniquely.

A tensor product is defined with the implied summation over repeated indices:

$$
A\, B \equiv {A_{\mu \nu}}^{\alpha \beta} {B_{\alpha \beta}}^{\rho \sigma}\,.
$$

**Magically, these tensor basis expressions form an algebra under this product, but I do not understand why. I can only verify that this fact indeed takes place.**

Let us correct one symmetric tensor product:

$$
X_{6} X_{6} = 8 X_{5} + 2 X_{6}.
$$

The definitions of the product and index symmetrization $(\cdot, \cdot)$ lead to $X_{6} X_{6}$ expressed as

$$
\begin{align}
& \equiv \frac{1}{k^2} k_{(\mu} {g_{\nu)}}^{(\alpha} k^{\beta)} \frac{1}{k^2} k_{(\alpha} {g_{\beta)}}^{(\rho} k^{\sigma)}\\
 &= \frac{1}{k^4} \Big( k_{\mu} {g_{\nu}}^{\alpha} k^{\beta} + k_{\mu} {g_{\nu}}^{\beta} k^{\alpha} \\
            & \;\;\;\;+ k_{\nu} {g_{\mu}}^{\alpha} k^{\beta} + k_{\nu} {g_{\mu}}^{\beta} k^{\alpha} \Big)\\
            & \;\;\;\;\cdot \Big( k_{\alpha} {g_{\beta}}^{\rho} k^{\sigma} + k_{\alpha} {g_{\beta}}^{\sigma} k^{\rho} \\
            & \;\;\;\;+ k_{\beta} {g_{\alpha}}^{\rho} k^{\sigma} + k_{\beta} {g_{\alpha}}^{\sigma} k^{\rho} \Big).
\end{align}
$$

This will produce 16 terms, but they will further simplify with the use of

$$
\begin{align}
k_{\alpha} k^{\alpha} & = k^2 \\
{g_{\nu}}^{\alpha}k_{\alpha} &= k_{\nu} \\
{g_{\nu}}^{\beta} {g_{\beta}}^{\sigma} &= {g_{\nu}}^{\sigma}\;.
\end{align}
$$

I will not write the whole process, but let us multiply the first terms in the parentheses, just to give one an idea:

$$
\begin{align}
t_{1} & = k_{\mu} {g_{\nu}}^{\alpha} k^{\beta} \cdot k_{\alpha} {g_{\beta}}^{\rho} k^{\sigma} \\
& = k_{\mu} \big({g_{\nu}}^{\alpha} k_{\alpha}\big) \big({g_{\beta}}^{\rho} k^{\beta}\big) k^{\sigma} \\
&= k_{\mu} k_{\nu} k^{\rho} k^{\sigma}.
\end{align}
$$

A page later ;) there will be 8 such terms and a doubled $X_{6}$ left. It is $8 X_{5} + 2 X_{6}$, not $8 X_{4} + 2 X_{6}$ in [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018).

All the computed traces of the algebra of the symmetric two-tensor field are correct. I will only note here the definition of a trace:

$$
\begin{align}
\textit{tr}\,{X_{\mu \nu}}^{\rho \sigma}  \equiv {X_{\mu \nu}}^{\mu \nu}.
\end{align}
$$

In order to calculate these tensorial traces one has to know the values of $g_{ij}$ and the mixed tensor

$$
\begin{align}
{g_{i}}^{j}={\delta_{i}}^{j}=\begin{cases}
0 & i \ne j \\
1 & i = j
\end{cases}\;\;.
\end{align}
$$

There is no need to know these values when getting the product tables.

The products and traces determine the spectrum. [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) applies [the Faddeev - LeVerrier algorithm](https://en.wikipedia.org/wiki/Faddeev%E2%80%93LeVerrier_algorithm), while [K.J. Barnes (1963)](https://spiral.imperial.ac.uk/bitstream/10044/1/13413/2/Barnes-KJ-1963-PhD-Thesis.pdf) goes for the matrix projection operators. Mysteriously, the eigenvalues will have multiplicities which can be derived independently from the tensor product of two $sl(2, \mathbb{C})$ algebras in the theory of the Lorentz group (Lorentz with "t"). The eigenvector equations lead to the Lorenz gauge condition (Lorenz without "t").

[Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) shows how one can get even eigenvalue multiplicities ("half integer spins") with [the gamma matrices](https://en.wikipedia.org/wiki/Gamma_matrices). The case with a vector and a spinor index leads to the rapid growth of complexity.

The algebra is constructed from $g \equiv \eta$, a four-vector $p_{\mu}$, and a four-vector-like $\gamma_{\mu}$ which also carries suppressed spinor indices. One arranges all the distinct products, et voilà. What is not trivial is that one needs to treat the matrix $\boldsymbol{p} \equiv p\llap{\unicode{x2215}}$ as an independent primitive which effectively doubles the number of basis elements from five to ten. $\boldsymbol{p}$ has no vector indices and hides spinorial indices implicitly.

How does one parse expressions such as $\gamma_{\mu}\gamma^{\nu} \boldsymbol{p}$? Given $\mu$ and $\nu$, it is a product of three 4x4 matrices. The matrix product is implicit.

The resulting phrase is likely erroneous: "... we find, together with their antiparticles, the following groups of particles: a quadruplet, and two doublets."

Weinberg's QFT, Vol. 1, pages 229-233, would indicate that [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) must have missed a triplet and its anti-triplet, but the arguments rely on group theory which is an [overwhelming detour](https://physics.stackexchange.com/questions/269234/1-2-0-representation-of-the-lorentz-group-so1-3).

The main strength of this work is that it provides an original way to treat irreducibility without group theory. A lot can also be automated. [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) uses advanced Faddeev-LeVerrier algorithm which allows to compute eigenvectors. It works even in the presence of multiple eigenvalues, see e.g. [Helmberg and Wagner (1993).](https://core.ac.uk/download/pdf/81192811.pdf)

The eigenvectors are linear combinations of the basis elements of the algebra that depend on the eigenvalue $\lambda$. By zeroing out the eigenvalues of the particles we want to remove from a field, we obtain operators which act on tensor fields as auxiliary conditions. They become PDEs with $k_{\mu} \rightarrow \partial_{\mu}$. What do we get with the eigenvectors of nonzero eigenvalues?

The latter should lead to free field (wave) equations, which is visible in the alternative formalism of [Joos-Weinberg](https://en.wikipedia.org/wiki/Joos%E2%80%93Weinberg_equation) clarified a lot by [Ahluwalia (1991)](https://arxiv.org/abs/nucl-th/9704066). Adjusted free-space Green functions of these wave PDEs serve as propagators in [Feynman diagram-based calculations](https://cerncourier.com/a/quantum-electrodynamics/). The EM interactions are often added via [$\partial_{\mu}\rightarrow \partial_{\mu}-ieA_{\mu}$](https://en.wikipedia.org/wiki/Rarita%E2%80%93Schwinger_equation), but this mechanism does not always work.

Steven Weinberg did not spend much time on consistent wave equations, but rather on consistent S-matrices, which is closer to nuclear reactions and scattering experiments. [Dharam Vir Ahluwalia](https://arxiv.org/a/ahluwalia_d_1.html), [Adam Gillard](https://ir.canterbury.ac.nz/items/81a8dc3d-d6f4-446a-8392-0c2f4f69f033), and others probe the Joos-Weinberg formalism in modern domains such as neutrinos and dark matter.

Ultimately, [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) provides only a "combine-split" mathematical puzzle game. As an input, we specify the number of spinorial and tensorial indices, and the tensorial symmetry or antisymmetry. In the output, we receive the list of "particle types" and the list of PDEs to remove each type. We control PCT and the number of field components by choosing "gamma-like" primitives involving Pauli matrices.

According to group theory, combining indices means taking "tensor products $(m,n)\otimes (k,l)$ of Lorentz irreps", clf. Weinberg's QFT, Vol. 1, pages 229-233:

$(0,0)$: A scalar.

$(\frac{1}{2},\frac{1}{2})$: A single four-vector index.

$(\frac{1}{2},0)\oplus (0,\frac{1}{2})$: A full single spinor index.

$(1,1)$: Two symmetric tensor indices.

$(1,0)\oplus (0,1)$: Two asymmetric tensor indices.

Scalars are not used by [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018), but they appear in the spectrum as singlets. In a symmetric two-index tensor case, to remove a singlet also means to make the tensor traceless.

Combining a vector and a spinor index implies taking the product $(\frac{1}{2},\frac{1}{2}) \otimes (\frac{1}{2},0)\oplus (0,\frac{1}{2})$
which splits into a vector and an irreducible component $(1,\frac{1}{2}) \oplus (\frac{1}{2},1)$, clf. Weinberg's QFT, Vol. 1, page 232. We can call the latter a "spin 3/2" field. In the Joos-Weinberg formalism, this is a different irreducible component $(\frac{3}{2},0) \oplus (0,\frac{3}{2})$. It could still be produced in Shankland's algebra, but not with one vector and one spinor index.

Why a single "spin" parameter? $(m,n)$ is not the Kronecker product $m\otimes n$ of "rotations", but Shankland's eigenvalue multiplicities follow this pattern, and the scalar value $m+m$ also describes the maximal "degrees of freedom of a particle". The Kronecker product splits into "spin multiplicities" $s=|m-n|,\ldots, m+n$, e.g.

$$
\begin{align}
(1,\frac{1}{2}) = \frac{1}{2}\oplus 1 \oplus \frac{3}{2}\,.
\end{align}
$$

The eigenvalue multiplicities will have $2s+1$ values. In Shankland's terminology, we obtain a doublet, a triplet, and a quadruplet from the equation above. However, this collection does not constitute three particles. It is one half of the irreducible "spin 3/2" particle $(1,\frac{1}{2}) \oplus (\frac{1}{2},1)$.

I am very vague here, being oblivious to complex vs real, Cartesian vs direct vs Kronecker products, groups vs algebras... The point here is only to show that group theory: (i) allows to verify Shankland's eigenvalue multiplicities, and (ii) indicates that not every multiple eigenvalue is a "separate particle". The whole group theory/QFT thing is beyond cumbersome. It is hilarious to find out that even Weinberg (!) uses the word **complicated** several times in his QFT Vol 3.

Although not of direct relevance, a human-readable group theory-free way to build invariant cost functions can be found in the works of P.A. Zhilin: [1](http://teormeh.net/Zhilin_New/pdf/Zhilin_Invariant_rus.pdf), [2](http://teormeh.net/Zhilin_New/pdf/Zhilin_Invariant_eng.pdf). It leads to a different kind of irreducibility and "particles".

One can find some mildly successful uses/hints of tensor algebras in [Phys. Rev. 106, 1345 (1957)](https://journals.aps.org/pr/abstract/10.1103/PhysRev.106.1345); [Nuovo Cimento, 43, 475 (1966)](https://link.springer.com/article/10.1007/BF02752873); [Nuovo Cimento 47, 145 (1967)](https://link.springer.com/article/10.1007/BF02818340); [Phys. Rev. 153, 1652 (1967)](https://journals.aps.org/pr/abstract/10.1103/PhysRev.153.1652); [Phys. Rev. 161, 1631 (1967)](https://journals.aps.org/pr/abstract/10.1103/PhysRev.161.1631); [Phys. Rev. D 8, 2650 (1973)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.8.2650); [Nuovo Cimento 28, 409 (1975)](https://inspirehep.net/literature/98459); [Phys. Lett. B 301 4 339 (1993)](https://arxiv.org/abs/hep-th/9212008); [Phys. Rev. C 64, 015203 (2001)](https://arxiv.org/abs/hep-ph/0103172); [Phys. Rev. D 64, 125013 (2001)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.64.125013); [Hadronic J. 26, 351 (2003)](https://www.imath.kiev.ua/~nikitin/PAPER26.pdf); [Phys. Rev. D 67, 085021 (2003)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.67.085021); [Phys. Rev. D 67, 125011 (2003)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.67.125011); [Nucl. Phys. B724, 453 (2005)](https://arxiv.org/abs/hep-th/0505255); [Phys. Rev. D 74, 084036 (2006)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.74.084036); [P. Cvitanović (2008)](https://birdtracks.eu/); [V. Monchiet and G. Bonnet (2010)](https://royalsocietypublishing.org/doi/10.1098/rspa.2010.0149); [Phys. Rev. D 97, 115043 (2018)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.97.115043); [SUGRA and CDC](https://news.stonybrook.edu/facultystaff/qa-with-breakthrough-prize-winner-peter-van-nieuwenhuizen/)...

Who wants to clean the mess?

<div style="text-align:center">
<a style="font-size: 1rem;" href="https://youtu.be/Y183gJQ9yCY?t=20">Sign the contract big boy...</a>
<a href="https://voi.id/en/memori/185165">
<img src="article2/IronMike-min-rs.png" style="display: block; margin: auto;" title="Iron Mike" width="50%">
</a>
</div>
