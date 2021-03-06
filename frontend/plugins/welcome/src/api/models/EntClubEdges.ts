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
    EntActivities,
    EntActivitiesFromJSON,
    EntActivitiesFromJSONTyped,
    EntActivitiesToJSON,
    EntClubBranch,
    EntClubBranchFromJSON,
    EntClubBranchFromJSONTyped,
    EntClubBranchToJSON,
    EntClubType,
    EntClubTypeFromJSON,
    EntClubTypeFromJSONTyped,
    EntClubTypeToJSON,
    EntClubapplication,
    EntClubapplicationFromJSON,
    EntClubapplicationFromJSONTyped,
    EntClubapplicationToJSON,
    EntComplaint,
    EntComplaintFromJSON,
    EntComplaintFromJSONTyped,
    EntComplaintToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClubEdges
 */
export interface EntClubEdges {
    /**
     * Activities holds the value of the activities edge.
     * @type {Array<EntActivities>}
     * @memberof EntClubEdges
     */
    activities?: Array<EntActivities>;
    /**
     * ClubToComplaint holds the value of the ClubToComplaint edge.
     * @type {Array<EntComplaint>}
     * @memberof EntClubEdges
     */
    clubToComplaint?: Array<EntComplaint>;
    /**
     * Clubapplication holds the value of the clubapplication edge.
     * @type {Array<EntClubapplication>}
     * @memberof EntClubEdges
     */
    clubapplication?: Array<EntClubapplication>;
    /**
     * 
     * @type {EntClubBranch}
     * @memberof EntClubEdges
     */
    clubbranch?: EntClubBranch;
    /**
     * 
     * @type {EntClubType}
     * @memberof EntClubEdges
     */
    clubtype?: EntClubType;
    /**
     * 
     * @type {EntUser}
     * @memberof EntClubEdges
     */
    user?: EntUser;
    /**
     * Userclub holds the value of the userclub edge.
     * @type {Array<EntUser>}
     * @memberof EntClubEdges
     */
    userclub?: Array<EntUser>;
}

export function EntClubEdgesFromJSON(json: any): EntClubEdges {
    return EntClubEdgesFromJSONTyped(json, false);
}

export function EntClubEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClubEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activities': !exists(json, 'activities') ? undefined : ((json['activities'] as Array<any>).map(EntActivitiesFromJSON)),
        'clubToComplaint': !exists(json, 'clubToComplaint') ? undefined : ((json['clubToComplaint'] as Array<any>).map(EntComplaintFromJSON)),
        'clubapplication': !exists(json, 'clubapplication') ? undefined : ((json['clubapplication'] as Array<any>).map(EntClubapplicationFromJSON)),
        'clubbranch': !exists(json, 'clubbranch') ? undefined : EntClubBranchFromJSON(json['clubbranch']),
        'clubtype': !exists(json, 'clubtype') ? undefined : EntClubTypeFromJSON(json['clubtype']),
        'user': !exists(json, 'user') ? undefined : EntUserFromJSON(json['user']),
        'userclub': !exists(json, 'userclub') ? undefined : ((json['userclub'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntClubEdgesToJSON(value?: EntClubEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activities': value.activities === undefined ? undefined : ((value.activities as Array<any>).map(EntActivitiesToJSON)),
        'clubToComplaint': value.clubToComplaint === undefined ? undefined : ((value.clubToComplaint as Array<any>).map(EntComplaintToJSON)),
        'clubapplication': value.clubapplication === undefined ? undefined : ((value.clubapplication as Array<any>).map(EntClubapplicationToJSON)),
        'clubbranch': EntClubBranchToJSON(value.clubbranch),
        'clubtype': EntClubTypeToJSON(value.clubtype),
        'user': EntUserToJSON(value.user),
        'userclub': value.userclub === undefined ? undefined : ((value.userclub as Array<any>).map(EntUserToJSON)),
    };
}


