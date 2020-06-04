import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of as observableOf } from 'rxjs';
import { catchError, mergeMap, map, filter } from 'rxjs/operators';
import { HttpStatus } from 'app/types/types';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';

import {
  CreateNotificationRule,
  CreateNotificationRuleSuccess,
  CreateNotificationRuleFailure,
  GetNotificationRulesSuccess,
  GetNotificationRulesFailure,
  DeleteNotificationRule,
  DeleteNotificationRuleSuccess,
  DeleteNotificationRuleFailure,
  NotificationRuleActionTypes
} from './notification_rule.action';

import { NotificationRuleRequests } from './notification_rule.requests';

@Injectable()
export class NotificationRuleEffects {
  constructor(
    private actions$: Actions,
    private requests: NotificationRuleRequests
  ) { }

  @Effect()
  getDestinations$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.GET_ALL),
      mergeMap(() =>
        this.requests.getNotificationRules().pipe(
          map(resp => new GetNotificationRulesSuccess(resp)),
          catchError((error: HttpErrorResponse) =>
            observableOf(new GetNotificationRulesFailure(error))))));

  @Effect()
  getDestinationsFailure$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.GET_ALL_FAILURE),
      map(({ payload }: GetNotificationRulesFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get notifications: ${msg || payload.error}`
        });
      }));

  @Effect()
  createNotificationRule$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.CREATE),
      mergeMap(({ payload, username, password }: CreateNotificationRule) =>
      this.requests.createNotificationRule( payload, username, password ).pipe(
        map(() => new CreateNotificationRuleSuccess(payload)),
        catchError((error: HttpErrorResponse) =>
          observableOf(new CreateNotificationRuleFailure(error))))));

  @Effect()
  createNotificationRuleSuccess$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.CREATE_SUCCESS),
      map(({ payload  }: CreateNotificationRuleSuccess) => new CreateNotification({
      type: Type.info,
      message: `Created notification ${payload.name}.`
    })));

  @Effect()
  createNotificationRuleFailure$ = this.actions$.pipe(
    ofType(NotificationRuleActionTypes.CREATE_FAILURE),
    filter(({ payload }: CreateNotificationRuleFailure) => payload.status !== HttpStatus.CONFLICT),
    map(({ payload }: CreateNotificationRuleFailure) => new CreateNotification({
        type: Type.error,
        message: `Could not create notification: ${payload.error.error || payload}.`
      })));
    
  @Effect()
  deleteNotificationRule$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.DELETE),
      mergeMap(({ payload: { rule } }: DeleteNotificationRule) =>
        this.requests.deleteNotificationRule(rule).pipe(
          map(() => new DeleteNotificationRuleSuccess({rule})),
          catchError((error: HttpErrorResponse) =>
            observableOf(new DeleteNotificationRuleFailure(error))))));

  @Effect()
  deleteNotificationRuleSuccess$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.DELETE_SUCCESS),
      map(({ payload: { rule } }: DeleteNotificationRuleSuccess) => {
        return new CreateNotification({
          type: Type.info,
          message: `Deleted rule ${rule.name}.`
        });
      }));

  @Effect()
  deleteNotificationRuleFailure$ = this.actions$.pipe(
      ofType(NotificationRuleActionTypes.DELETE_FAILURE),
      map(({ payload }: DeleteNotificationRuleFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not delete rule: ${msg || payload.error}`
        });
      }));

}
