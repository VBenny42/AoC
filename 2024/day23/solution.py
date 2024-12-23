import networkx as nx


def main1():
    G = nx.Graph()
    with open("input.txt", "r", encoding="utf-8") as f:
        for line in f:
            u, v = line.strip().split("-")
            G.add_edge(u, v)

    triangles = []
    for clique in nx.enumerate_all_cliques(G):
        if len(clique) == 3:
            triangles.append(clique)
        if len(clique) > 3:
            break

    # Filter for cliques containing at least one node starting with 't'
    triangles_with_t = sum(
        1 for triangle in triangles if any(node.startswith("t") for node in triangle)
    )

    # Print results
    print(f"ANSWER1: {(triangles_with_t)}")


def main2():
    G = nx.Graph()
    with open("input.txt", "r", encoding="utf-8") as f:
        for line in f:
            u, v = line.strip().split("-")
            G.add_edge(u, v)

    max_clique = max(nx.find_cliques(G), key=len)
    max_clique = ",".join(sorted(max_clique))
    print(f"ANSWER2: { max_clique }")


if __name__ == "__main__":
    main1()
    main2()
