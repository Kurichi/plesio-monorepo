import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/provider/repository/tree_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'init_tree_usecase.g.dart';

@riverpod
InitTreeUseCase initTreeUseCase(InitTreeUseCaseRef ref) {
  final treeRepository = ref.watch(treeRepositoryProvider);
  return InitTreeUseCase(treeRepository);
}

class InitTreeUseCase {
  final TreeRepository _treeRepository;

  InitTreeUseCase(this._treeRepository);

  Future<void> execute() async {
    await _treeRepository.initTree();
  }
}
