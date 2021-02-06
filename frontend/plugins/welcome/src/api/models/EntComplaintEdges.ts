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
    EntComplaintType,
    EntComplaintTypeFromJSON,
    EntComplaintTypeFromJSONTyped,
    EntComplaintTypeToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntComplaintEdges
 */
export interface EntComplaintEdges {
    /**
     * 
     * @type {EntClub}
     * @memberof EntComplaintEdges
     */
    complaintToClub?: EntClub;
    /**
     * 
     * @type {EntComplaintType}
     * @memberof EntComplaintEdges
     */
    complaintToComplaintType?: EntComplaintType;
    /**
     * 
     * @type {EntUser}
     * @memberof EntComplaintEdges
     */
    complaintToUser?: EntUser;
}

export function EntComplaintEdgesFromJSON(json: any): EntComplaintEdges {
    return EntComplaintEdgesFromJSONTyped(json, false);
}

export function EntComplaintEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntComplaintEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'complaintToClub': !exists(json, 'ComplaintToClub') ? undefined : EntClubFromJSON(json['ComplaintToClub']),
        'complaintToComplaintType': !exists(json, 'ComplaintToComplaintType') ? undefined : EntComplaintTypeFromJSON(json['ComplaintToComplaintType']),
        'complaintToUser': !exists(json, 'ComplaintToUser') ? undefined : EntUserFromJSON(json['ComplaintToUser']),
    };
}

export function EntComplaintEdgesToJSON(value?: EntComplaintEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'complaintToClub': EntClubToJSON(value.complaintToClub),
        'complaintToComplaintType': EntComplaintTypeToJSON(value.complaintToComplaintType),
        'complaintToUser': EntUserToJSON(value.complaintToUser),
    };
}


