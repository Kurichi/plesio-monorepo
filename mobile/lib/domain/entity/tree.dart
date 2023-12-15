class Tree {
  final String id;
  int growthLevel;

  Tree({
    required this.id,
    required this.growthLevel,
  });

  void grow(int growthAmount) {
    growthLevel += growthAmount;
  }
}
