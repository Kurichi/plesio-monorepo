import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';

class GetMyTreeUseCase {
  final TreeRepository _treeRepository;

  GetMyTreeUseCase(this._treeRepository);

  Future<Tree> execute() async {
    return await _treeRepository.getMyTree();
  }
}
