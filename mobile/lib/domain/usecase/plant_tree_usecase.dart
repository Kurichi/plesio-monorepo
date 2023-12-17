import 'package:kiikuten/domain/repository/tree_repository.dart';

class PlantTreeUseCase {
  final TreeRepository _treeRepository;

  PlantTreeUseCase(this._treeRepository);

  Future<void> execute() async {
    await _treeRepository.plantTree();
  }
}
