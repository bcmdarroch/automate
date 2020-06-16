import { desktopEntityReducer, DesktopEntityState, desktopEntityInitialState } from './desktop.reducer';
import { UpdateDesktopsFilterPageSizeAndCurrentPage } from './desktop.actions';

describe('desktopEntityReducer', () => {
  const initialState: DesktopEntityState = desktopEntityInitialState;

  describe('DesktopActionTypes', () => {
    describe('UPDATE_DESKTOPS_FILTER_PAGE_SIZE_AND_CURRENT_PAGE', () => {

      it('updates the pageSize and pageNumber', () => {
        const action = new UpdateDesktopsFilterPageSizeAndCurrentPage({
          pageSize: 20, updatedPageNumber: 3
        });
        const newState: any = desktopEntityReducer(initialState, action);

        expect(newState.getDesktopsFilter.pageSize).toBe(20);
        expect(newState.getDesktopsFilter.currentPage).toBe(3);
      });
    });
  });
});
