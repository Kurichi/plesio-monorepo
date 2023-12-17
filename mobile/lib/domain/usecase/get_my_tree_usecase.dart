import 'package:kiikuten/domain/entity/tree.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/provider/repository/tree_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'get_my_tree_usecase.g.dart';

@riverpod
GetMyTreeUseCase getMyTreeUseCase(GetMyTreeUseCaseRef ref) {
  final treeRepository = ref.watch(treeRepositoryProvider);
  return GetMyTreeUseCase(treeRepository);
}

class GetMyTreeUseCase {
  final TreeRepository _treeRepository;

  GetMyTreeUseCase(this._treeRepository);

  Future<Tree> execute() async {
    return await _treeRepository.getMyTree();
  }
}
