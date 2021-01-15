/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntClub,
    EntClubFromJSON,
    EntClubFromJSONTyped,
    EntClubToJSON,
    EntClubappStatus,
    EntClubappStatusFromJSON,
    EntClubappStatusFromJSONTyped,
    EntClubappStatusToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClubapplicationEdges
 */
export interface EntClubapplicationEdges {
    /**
     * 
     * @type {EntClub}
     * @memberof EntClubapplicationEdges
     */
    club?: EntClub;
    /**
     * 
     * @type {EntClubappStatus}
     * @memberof EntClubapplicationEdges
     */
    clubappstatus?: EntClubappStatus;
    /**
     * 
     * @type {EntUser}
     * @memberof EntClubapplicationEdges
     */
    owner?: EntUser;
}

export function EntClubapplicationEdgesFromJSON(json: any): EntClubapplicationEdges {
    return EntClubapplicationEdgesFromJSONTyped(json, false);
}

export function EntClubapplicationEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClubapplicationEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'club': !exists(json, 'club') ? undefined : EntClubFromJSON(json['club']),
        'clubappstatus': !exists(json, 'clubappstatus') ? undefined : EntClubappStatusFromJSON(json['clubappstatus']),
        'owner': !exists(json, 'owner') ? undefined : EntUserFromJSON(json['owner']),
    };
}

export function EntClubapplicationEdgesToJSON(value?: EntClubapplicationEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'club': EntClubToJSON(value.club),
        'clubappstatus': EntClubappStatusToJSON(value.clubappstatus),
        'owner': EntUserToJSON(value.owner),
    };
}

