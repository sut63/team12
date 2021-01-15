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
    EntClubapplication,
    EntClubapplicationFromJSON,
    EntClubapplicationFromJSONTyped,
    EntClubapplicationToJSON,
    EntComplaint,
    EntComplaintFromJSON,
    EntComplaintFromJSONTyped,
    EntComplaintToJSON,
    EntDiscipline,
    EntDisciplineFromJSON,
    EntDisciplineFromJSONTyped,
    EntDisciplineToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntRoomuse,
    EntRoomuseFromJSON,
    EntRoomuseFromJSONTyped,
    EntRoomuseToJSON,
    EntUserStatus,
    EntUserStatusFromJSON,
    EntUserStatusFromJSONTyped,
    EntUserStatusToJSON,
    EntUsertype,
    EntUsertypeFromJSON,
    EntUsertypeFromJSONTyped,
    EntUsertypeToJSON,
    EntYear,
    EntYearFromJSON,
    EntYearFromJSONTyped,
    EntYearToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * Club holds the value of the club edge.
     * @type {Array<EntClub>}
     * @memberof EntUserEdges
     */
    club?: Array<EntClub>;
    /**
     * Clubapplication holds the value of the clubapplication edge.
     * @type {Array<EntClubapplication>}
     * @memberof EntUserEdges
     */
    clubapplication?: Array<EntClubapplication>;
    /**
     * 
     * @type {EntDiscipline}
     * @memberof EntUserEdges
     */
    discipline?: EntDiscipline;
    /**
     * 
     * @type {EntClub}
     * @memberof EntUserEdges
     */
    fromClub?: EntClub;
    /**
     * 
     * @type {EntGender}
     * @memberof EntUserEdges
     */
    gender?: EntGender;
    /**
     * Roomuse holds the value of the Roomuse edge.
     * @type {Array<EntRoomuse>}
     * @memberof EntUserEdges
     */
    roomuse?: Array<EntRoomuse>;
    /**
     * UserToComplaint holds the value of the UserToComplaint edge.
     * @type {Array<EntComplaint>}
     * @memberof EntUserEdges
     */
    userToComplaint?: Array<EntComplaint>;
    /**
     * 
     * @type {EntUserStatus}
     * @memberof EntUserEdges
     */
    userstatus?: EntUserStatus;
    /**
     * 
     * @type {EntUsertype}
     * @memberof EntUserEdges
     */
    usertype?: EntUsertype;
    /**
     * 
     * @type {EntYear}
     * @memberof EntUserEdges
     */
    year?: EntYear;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'club': !exists(json, 'Club') ? undefined : ((json['Club'] as Array<any>).map(EntClubFromJSON)),
        'clubapplication': !exists(json, 'Clubapplication') ? undefined : ((json['Clubapplication'] as Array<any>).map(EntClubapplicationFromJSON)),
        'discipline': !exists(json, 'Discipline') ? undefined : EntDisciplineFromJSON(json['Discipline']),
        'fromClub': !exists(json, 'FromClub') ? undefined : EntClubFromJSON(json['FromClub']),
        'gender': !exists(json, 'Gender') ? undefined : EntGenderFromJSON(json['Gender']),
        'roomuse': !exists(json, 'Roomuse') ? undefined : ((json['Roomuse'] as Array<any>).map(EntRoomuseFromJSON)),
        'userToComplaint': !exists(json, 'UserToComplaint') ? undefined : ((json['UserToComplaint'] as Array<any>).map(EntComplaintFromJSON)),
        'userstatus': !exists(json, 'Userstatus') ? undefined : EntUserStatusFromJSON(json['Userstatus']),
        'usertype': !exists(json, 'Usertype') ? undefined : EntUsertypeFromJSON(json['Usertype']),
        'year': !exists(json, 'Year') ? undefined : EntYearFromJSON(json['Year']),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'club': value.club === undefined ? undefined : ((value.club as Array<any>).map(EntClubToJSON)),
        'clubapplication': value.clubapplication === undefined ? undefined : ((value.clubapplication as Array<any>).map(EntClubapplicationToJSON)),
        'discipline': EntDisciplineToJSON(value.discipline),
        'fromClub': EntClubToJSON(value.fromClub),
        'gender': EntGenderToJSON(value.gender),
        'roomuse': value.roomuse === undefined ? undefined : ((value.roomuse as Array<any>).map(EntRoomuseToJSON)),
        'userToComplaint': value.userToComplaint === undefined ? undefined : ((value.userToComplaint as Array<any>).map(EntComplaintToJSON)),
        'userstatus': EntUserStatusToJSON(value.userstatus),
        'usertype': EntUsertypeToJSON(value.usertype),
        'year': EntYearToJSON(value.year),
    };
}


