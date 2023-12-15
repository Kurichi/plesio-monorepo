class Tree {
  final String id;
  final String type;
  int growthLevel;

  Tree({
    required this.id,
    required this.type,
    required this.growthLevel,
  });

  void grow(int growthAmount) {
    growthLevel += growthAmount;
  }
}
