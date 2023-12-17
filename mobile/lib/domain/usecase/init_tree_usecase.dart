import 'package:kiikuten/domain/repository/tree_repository.dart';

class InitTreeUseCase {
  final TreeRepository _treeRepository;

  InitTreeUseCase(this._treeRepository);

  Future<void> execute() async {
    await _treeRepository.initTree();
  }
}
