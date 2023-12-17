import 'package:kiikuten/domain/entity/tree_ranking.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/provider/repository/tree_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'get_tree_ranking_usecase.g.dart';

@riverpod
GetTreeRankingUseCase getTreeRankingUseCase(GetTreeRankingUseCaseRef ref) {
  final treeRepository = ref.watch(treeRepositoryProvider);
  return GetTreeRankingUseCase(treeRepository);
}

class GetTreeRankingUseCase {
  final TreeRepository _treeRepository;

  GetTreeRankingUseCase(this._treeRepository);

  Future<List<TreeRanking>> execute() async {
    return await _treeRepository.getTreeRanking();
  }
}
