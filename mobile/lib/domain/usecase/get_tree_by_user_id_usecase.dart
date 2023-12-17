import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/provider/repository/tree_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'get_tree_by_user_id_usecase.g.dart';

@riverpod
GetTreeByUserIdUseCase getTreeByUserIdUseCase(GetTreeByUserIdUseCaseRef ref) {
  final treeRepository = ref.watch(treeRepositoryProvider);
  return GetTreeByUserIdUseCase(treeRepository);
}

class GetTreeByUserIdUseCase {
  final TreeRepository _treeRepository;

  GetTreeByUserIdUseCase(this._treeRepository);

  Future<Tree> execute(String userId) async {
    return await _treeRepository.getTreeByUserId(userId);
  }
}
