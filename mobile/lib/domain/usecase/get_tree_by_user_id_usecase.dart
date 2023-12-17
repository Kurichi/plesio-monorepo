import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';

class GetTreeByUserIdUseCase {
  final TreeRepository _treeRepository;

  GetTreeByUserIdUseCase(this._treeRepository);

  Future<Tree> execute(String userId) async {
    return await _treeRepository.getTreeByUserId(userId);
  }
}
